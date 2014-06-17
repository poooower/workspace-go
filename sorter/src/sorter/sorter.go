package sorter

import (
	"algorithms"
	"algorithms/bubblesort"
	"algorithms/qsort"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Sorter struct {
	infile    string
	outfile   string
	algorithm string
}

func NewSorter(infile string, outfile string, algorithm string) Sorter {
	return Sorter{infile, outfile, algorithm}
}

func (this *Sorter) DoSort() {
	if len(this.infile) > 0 {
		fmt.Println("hello infile =", this.infile, "outfile =", this.outfile, "algorithm =", this.algorithm)
		// fmt.Println("hello world")
	}

	values, err := this.readValues(this.infile)
	if err == nil {
		fmt.Println("Read values:", values)

		var sort algorithms.Sort
		switch this.algorithm {
		case "bubblesort":
			sort = bubblesort.NewBubbleSort()
		case "qsort":
			sort = qsort.NewQuickSort()
		}
		if sort == nil {
			fmt.Println("algorithm=", this.algorithm, " is not support")
			return
		}

		t1 := time.Now()
		sort.Sort(values)
		t2 := time.Now()

		fmt.Println("The", this.algorithm, "sorting process costs", t2.Sub(t1), "to complete.")
		this.writeValues(values, this.outfile)

	} else {
		fmt.Println(err)
	}
}

func (this *Sorter) readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		str := string(line) // 转换字符数组为字符串
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func (this *Sorter) writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()
	file.WriteString(this.algorithm + " result:\n")
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
