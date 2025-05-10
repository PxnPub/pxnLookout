package protocols;

import(
	Log     "log"
	Fmt     "fmt"
	Time    "time"
	Sort    "sort"
	Strings "strings"
	StrConv "strconv"
	TamDB   "github.com/PxnPub/TamDB"
	GoSNMP  "github.com/gosnmp/gosnmp"
	NatSort "github.com/maruel/natural"
	Configs "github.com/PxnPub/pxnLookout/LookoutBackend/configs"
);



const SNMP_OID_IFOCTIN  = "1.3.6.1.2.1.31.1.1.1.6.%d";
const SNMP_OID_IFOCTOUT = "1.3.6.1.2.1.31.1.1.1.10.%d";



type MonProtocol_SNMP struct {
	Name     string
	Interval *Time.Duration
	Tam      *TamDB.TamDB
	Snmp     *GoSNMP.GoSNMP
	Oids     []string
	Nodes    []string
	Fields   []string
	States   map[string]int64
}

type SNMP_Node struct {
	Node  string
	Field string
}



func NewMonProtocol_SNMP(name string, cfg *Configs.ConfigMonProto_SNMP,
		interval *Time.Duration, tam *TamDB.TamDB) (*MonProtocol_SNMP, error) {
	// connect snmp
	snmp := GoSNMP.GoSNMP{
		Target:    cfg.Host,
		Port:      161,
		Community: cfg.Community,
		Version:   GoSNMP.Version2c,
		Timeout:   Time.Duration(5 * Time.Second),
	};
	if err := snmp.Connect(); err != nil { return nil, err; }
	// nodes/fields
	nodes  := make([]string, 0);
	nods   := make([]string, 0);
	oids   := make([]string, 0);
	fields := make([]string, 0);
	for node, _ := range cfg.Nodes {
		nodes = append(nodes, node);
	}
	Sort.Sort(NatSort.StringSlice(nodes));
	for _, node := range nodes {
		field := cfg.Nodes[node];
		switch {
			// ethernet bandwidth
			case Strings.HasPrefix(node, "eth-"): {
				idx, err := StrConv.Atoi(Strings.TrimPrefix(node, "eth-"));
				if err != nil { return nil, err; }
				field_in  := Fmt.Sprintf("%s-in",  field);
				field_out := Fmt.Sprintf("%s-out", field);
				oids   = append(oids, Fmt.Sprintf(SNMP_OID_IFOCTIN,  idx));
				oids   = append(oids, Fmt.Sprintf(SNMP_OID_IFOCTOUT, idx));
				nods   = append(nods, Fmt.Sprintf("%s-in",  node));
				nods   = append(nods, Fmt.Sprintf("%s-out", node));
				fields = append(fields, field_in);
				fields = append(fields, field_out);
				tam.AddField(field_in);
				tam.AddField(field_out);
				Fmt.Printf(
					"     Node: %s  ->  %d %s  %d %s\n",
					node, len(oids)-2, field_in, len(oids)-1, field_out,
				);
			}
			default: return nil, Fmt.Errorf("Invalid SNMP node: %s", node);
		}
	}
	return &MonProtocol_SNMP{
		Name:     name,
		Interval: interval,
		Tam:      tam,
		Snmp:     &snmp,
		Oids:     oids,
		Nodes:    nods,
		Fields:   fields,
		States:   make(map[string]int64),
	}, nil;
}

func (proto *MonProtocol_SNMP) Close() {
	proto.Snmp.Conn.Close();
	proto.Tam.Close();
}



func (proto *MonProtocol_SNMP) Loop(timestamp int64) {
	result, err := proto.Snmp.Get(proto.Oids);
	if err != nil { Log.Panic(err); }
	count := len(proto.Fields);
	for index:=0; index<count; index++ {
		node  := proto.Nodes[index];
		field := proto.Fields[index];
		switch {
			// ethernet bandwidth
			case Strings.HasPrefix(node, "eth-"): {
				val := int64(result.Variables[index].Value.(uint64));
				state, exists := proto.States[field];
				if exists {
					bw := float64(val - state) / (float64(*proto.Interval) / float64(Time.Second));
					proto.Tam.Submit(timestamp, field, int64(bw));
				}
				proto.States[field] = val;
			}
		}
	}
}
