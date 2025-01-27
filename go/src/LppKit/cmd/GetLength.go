// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"bytes"

	"lpp"
	"os"

	"github.com/spf13/cobra"
)

// GetLengthCmd represents the GetLength command
var GetLengthCmd = &cobra.Command{
	Use:   "GetLength",
	Short: "Get length in a fasta file.",
	Long:  `Get length in a fasta file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fasta := getFlagString(cmd, "fasta")

		output := getFlagString(cmd, "output")
		OUTPUTHANDLE, err := os.Create(*output)
		defer OUTPUTHANDLE.Close()
		OUTPUTBUF := bufio.NewWriterSize(OUTPUTHANDLE, 10000)
		OUTPUTBUF.WriteString("Name\tLength\n")
		defer OUTPUTBUF.Flush()
		if err != nil {
			panic("Output not Exist!!")
		}

		FASTAIO := lpp.GetBlockRead(*fasta, "\n>", false, 10000000)
		for {
			line, err := FASTAIO.Next()

			line = bytes.TrimSuffix(line, []byte(">"))
			line = bytes.TrimPrefix(line, []byte(">"))
			name := bytes.Fields(line)[0]
			seq := bytes.SplitN(line, []byte("\n"), 2)[1]
			seq = bytes.Replace(seq, []byte("\n"), []byte(""), -1)
			length := len(seq)

			output_byte := name
			output_byte = append(output_byte, []byte("\t")...)
			output_byte = append(output_byte, []byte(string(length))...)
			output_byte = append(output_byte, []byte("\n")...)
			OUTPUTBUF.Write(output_byte)
			if err != nil {
				break
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(GetLengthCmd)
	GetLengthCmd.PersistentFlags().StringP("fasta", "f", "", "Fasta Input")
	GetLengthCmd.PersistentFlags().StringP("output", "o", "", "Fasta Output")
	GetLengthCmd.MarkFlagRequired("output")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GetLengthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GetLengthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
