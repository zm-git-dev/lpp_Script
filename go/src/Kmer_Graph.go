package main

import (
	"bytes"
	"flag"
	//"fmt"
	"lpp"
	"strconv"
	"os"
	"regexp"
	"strings"
)

var x int

var path []string
var kmer_graph map[string]map[string]string = make(map[string]map[string]string)
var kmer_seq map[string]string = make(map[string]string)

var OUTPUT *os.File

func Contains(s_list []string, node string) bool {
	res := false
	for _, data := range s_list {
		if data == node {
			res = true
		}
	}
	return res
}
func Traverse_5(order int, node string, step int, path []string, threshold int) []string {
	if Contains(path, node) {
		return path
	}
	path = append(path, node)
	//	_, ok := kmer_graph[node]

	if step == threshold {
		if len(path) > 0 {
			OUTPUT.WriteString(">" + strings.Join(path, "; ") + "\n")
			end_seq := ""
			for i, name := range path {
				new_seq, _ := kmer_seq[name]
				if i == 0 {

					end_seq = new_seq
				} else {
					end_seq += new_seq[*kmer-1:]
				}
			}
			if order%2 == 1 {
				end_seq = string(lpp.RevComplement([]byte(end_seq)))
			}
			OUTPUT.WriteString(end_seq + "\n")
		}

	} else {
		loc := Find(path, node)
		if len(kmer_graph[node]) > 1 {
			step += 1
		}
		for son, _ := range kmer_graph[node] {
			if len(kmer_graph[node]) > 1 {
				path = path[:loc]
			}

			path = Traverse_5(order, son, step, path, threshold)
		}

	}
	return path
}
func Find(s_list []string, node string) int {
	j := 0
	for i, data := range s_list {
		if data == node {
			j = i
			break
		}
	}
	return j + 1
}

var kmer *int

func main() {
	has_son := make(map[string]string)
	input := flag.String("i", "./", "Input File")
	output := flag.String("o", "StringGraph.tsv", "Output File")
	nodes_file := flag.String("l", "KmerList.tsv", "Kmer Id list")
	kmer = flag.Int("k", 41, "kmer")
	threshold := flag.Int("s", 10, "Step")
	flag.Parse()
	raw_file := lpp.GetBlockRead(*nodes_file, "\n", false, 10000)
	var All_list  []string
	for {
		line,err := raw_file.Next()
		line = bytes.TrimSpace(line)
		All_list = append( All_list, string(line)  )
		if err!=nil{
			break
		}

	}
	RAW := lpp.Fasta{File: *input}
	OUTPUT, _ = lpp.GetOuput(*output, 1000)

	reg := regexp.MustCompile(`L\:(\S)\:(\d+)\:(\S)`)
	reg_cov := regexp.MustCompile(`KC\:i\:(\d+)`)

	for {
		title, seq, err := RAW.Next()
		name := string(bytes.Fields(title)[0][1:])
		all_situation := reg.FindAllStringSubmatch(string(title), -1)
		cov,_ := strconv.Atoi( reg_cov.FindStringSubmatch(string(title))[1] )
		if cov<30{
			continue
		}
		if len(all_situation) == 0 {
			continue
		}
		if len(all_situation) < 2 {
			continue
		}
		has_son[name] = ""
		seq = bytes.TrimSpace(seq)
		kmer_seq[name+"+"] = string(seq)
		kmer_seq[name+"-"] = string(lpp.RevComplement(seq))
		if err != nil {
			break
		}
	}
	RAW = lpp.Fasta{File: *input}
	for {
		title, _, err := RAW.Next()
		name := string(bytes.Fields(title)[0][1:])
		cov,_ := strconv.Atoi( reg_cov.FindStringSubmatch(string(title))[1] )
		if cov<30{
			continue
		}
		_, ok := has_son[name]
		if !ok {
			continue
		}
		all_situation := reg.FindAllStringSubmatch(string(title), -1)
		for _, data := range all_situation {
			dir := data[1]
			sub := data[2]
			_, ok := has_son[sub]
			if !ok {
				continue
			}
			sub_dir := data[3]
			q := name + dir
			s := sub + sub_dir
			_, ok1 := kmer_graph[q]

			if !ok1 {
				kmer_graph[q] = make(map[string]string)

			}
			kmer_graph[q][s] = ""
			if dir == "+" {
				dir = "-"
			} else {
				dir = "+"
			}
			if sub_dir == "+" {
				sub_dir = "-"
			} else {
				sub_dir = "+"
			}
			q = sub + sub_dir
			s = name + dir
			_, ok2 := kmer_graph[q]
			if !ok2 {
				kmer_graph[q] = make(map[string]string)

			}
			kmer_graph[q][s] = ""

		}
		if err != nil {
			break
		}
	}
	for n, node := range All_list {
		OUTPUT.WriteString(node + "\n")
		Traverse_5(n, node, 0, path, *threshold)
		

	}
}
