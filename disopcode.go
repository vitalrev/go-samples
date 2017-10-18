package main

import (
	"bytes"
	"regexp"
	"strings"
	"fmt"
	"strconv"
)

func disasm(code string) string {
	if code == "" {
		return code
	}
	re := regexp.MustCompile("(..?)")
	//codes := code.match(/(..?)/g);  //JS
	codes := re.FindAllString(code, -1)
	//fmt.Println(codes)
	//var dis string;
	var dis bytes.Buffer
	for i := 1; i < len(codes); i++ {
		opcode, ok := opcodes[codes[i]]
		//fmt.Println("check: ", codes[i], opcode)
		if !ok {
			dis.WriteString("Missing opcode \n")
		} else if len(opcode) > 4 && opcode[0:4] == "PUSH" {
			var s string = strings.Replace(opcode, "PUSH", "", -1)
			//fmt.Printf("replaced PUSH: %s\n", s)
			length, _ := strconv.Atoi(s)
			dis.WriteString(opcode)
			dis.WriteString(" 0x")
			var s1 []string = codes[(i+1) : (i+length+1)]
			dis.WriteString(strings.Join(s1, ""))
			dis.WriteString("\n")
			i = i + length
		} else {
			dis.WriteString(opcode)
			dis.WriteString("\n")
		}
	}

	return dis.String()
}

func main() {
	var bytecode string = "0x60606040523615610055576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063661172761461014857806382c90ac01461017e578063b76ea962146101b4575b6101465b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163460405180905060006040518083038185876187965a03f1925050501561013d577f23919512b2162ddc59b67a65e3b03c419d4105366f7d4a632f5d3c3bee9b1cff600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1610143565b60006000fd5b5b565b005b341561015057fe5b61017c600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610225565b005b341561018657fe5b6101b2600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506102c8565b005b610223600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190505061036b565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156102825760006000fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b50565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103255760006000fd5b80600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b50565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103c85760006000fd5b8173ffffffffffffffffffffffffffffffffffffffff1634826040518082805190602001908083836000831461041d575b80518252602083111561041d576020820191506020810190506020830392506103f9565b505050905090810190601f1680156104495780820380516001836020036101000a031916815260200191505b5091505060006040518083038185876187965a03f192505050151561046e5760006000fd5b5b5b50505600a165627a7a7230582015dd8a30edc4d2b88e51c1e252cdc5a08cd0159c2d5d7b063fdaad85d6e813e40029"

	opcode := disasm(bytecode)
	fmt.Println(opcode)
}

var opcodes = map[string]string{
	"00": "STOP",
	"01": "ADD",
	"02": "MUL",
	"03": "SUB",
	"04": "DIV",
	"05": "SDIV",
	"06": "MOD",
	"07": "SMOD",
	"08": "ADDMOD",
	"09": "MULMOD",
	"0a": "EXP",
	"0b": "SIGNEXTEND",
	"10": "LT",
	"11": "GT",
	"12": "SLT",
	"13": "SGT",
	"14": "EQ",
	"15": "ISZERO",
	"16": "AND",
	"17": "OR",
	"18": "XOR",
	"19": "NOT",
	"1a": "BYTE",
	"20": "SHA3",
	"30": "ADDRESS",
	"31": "BALANCE",
	"32": "ORIGIN",
	"33": "CALLER",
	"34": "CALLVALUE",
	"35": "CALLDATALOAD",
	"36": "CALLDATASIZE",
	"37": "CALLDATACOPY",
	"38": "CODESIZE",
	"39": "CODECOPY",
	"3a": "GASPRICE",
	"3b": "EXTCODESIZE",
	"3c": "EXTCODECOPY",
	"40": "BLOCKHASH",
	"41": "COINBASE",
	"42": "TIMESTAMP",
	"43": "NUMBER",
	"44": "DIFFICULTY",
	"45": "GASLIMIT",
	"50": "POP",
	"51": "MLOAD",
	"52": "MSTORE",
	"53": "MSTORE8",
	"54": "SLOAD",
	"55": "SSTORE",
	"56": "JUMP",
	"57": "JUMPI",
	"58": "PC",
	"59": "MSIZE",
	"5a": "GAS",
	"5b": "JUMPDEST",
	"60": "PUSH1",
	"61": "PUSH2",
	"62": "PUSH3",
	"63": "PUSH4",
	"64": "PUSH5",
	"65": "PUSH6",
	"66": "PUSH7",
	"67": "PUSH8",
	"68": "PUSH9",
	"69": "PUSH10",
	"6a": "PUSH11",
	"6b": "PUSH12",
	"6c": "PUSH13",
	"6d": "PUSH14",
	"6e": "PUSH15",
	"6f": "PUSH16",
	"70": "PUSH17",
	"71": "PUSH18",
	"72": "PUSH19",
	"73": "PUSH20",
	"74": "PUSH21",
	"75": "PUSH22",
	"76": "PUSH23",
	"77": "PUSH24",
	"78": "PUSH25",
	"79": "PUSH26",
	"7a": "PUSH27",
	"7b": "PUSH28",
	"7c": "PUSH29",
	"7d": "PUSH30",
	"7e": "PUSH31",
	"7f": "PUSH32",
	"80": "DUP1",
	"81": "DUP2",
	"82": "DUP3",
	"83": "DUP4",
	"84": "DUP5",
	"85": "DUP6",
	"86": "DUP7",
	"87": "DUP8",
	"88": "DUP9",
	"89": "DUP10",
	"8a": "DUP11",
	"8b": "DUP12",
	"8c": "DUP13",
	"8d": "DUP14",
	"8e": "DUP15",
	"8f": "DUP16",
	"90": "SWAP1",
	"91": "SWAP2",
	"92": "SWAP3",
	"93": "SWAP4",
	"94": "SWAP5",
	"95": "SWAP6",
	"96": "SWAP7",
	"97": "SWAP8",
	"98": "SWAP9",
	"99": "SWAP10",
	"9a": "SWAP11",
	"9b": "SWAP12",
	"9c": "SWAP13",
	"9d": "SWAP14",
	"9e": "SWAP15",
	"9f": "SWAP16",
	"a0": "LOG0",
	"a1": "LOG1",
	"a2": "LOG2",
	"a3": "LOG3",
	"f0": "CREATE",
	"f1": "CALL",
	"f2": "CALLCODE",
	"f3": "RETURN",
	"ff": "SUICIDE",
}
