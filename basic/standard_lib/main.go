package main

import (
	"bufio"
	"container/list"
	"container/ring"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func main() {
	runErrors()
	runFlag()

	fmt.Println(strings.Contains("Kharisma", "aris"))
	fmt.Println(strings.Split("Kharisma", "aris"))
	fmt.Println(strings.ToLower("Kharisma"))
	fmt.Println(strings.ToUpper("kharisma"))
	fmt.Println(strings.Trim("Kharisma ", " "))
	fmt.Println(strings.ReplaceAll("Kharisma", "ma", "wa"))

	var a int = 20
	fmt.Println(strconv.Itoa(a))

	// angka := "A"
	angka := "20"
	result, err := strconv.Atoi(angka)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

	fmt.Println(math.Ceil(1.60))  // bulat ke atas
	fmt.Println(math.Floor(1.60)) // bulat ke bawah
	fmt.Println(math.Round(1.60)) // bulat ke terdekat
	fmt.Println(math.Max(1, 60))
	fmt.Println(math.Min(1, 60))

	// var data *list.List = list.New()
	data := list.New()
	data.PushBack("Kharis")
	data.PushBack("Nanda")
	data.PushBack("Wardhana")
	data.PushFront("Nama")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// var dataCircular *ring.Ring = ring.New(5)
	dataCircular := ring.New(5)
	for i := 0; i < dataCircular.Len(); i++ {
		dataCircular.Value = i
		dataCircular.Next()
	}
	// read dataCircular
	dataCircular.Do(func(val any) {
		fmt.Println(val)
	})

	// sort
	users := []User{
		{"Kharis", 30},
		{"Andit", 20},
		{"Badi", 28},
	}
	sort.Sort(UserSlice(users))
	fmt.Println(users)

	fmt.Println(time.Now())
	utc := time.Date(2024, time.January, 22, 13, 50, 00, 00, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	// yyyy-MM-dd HH:mm:ss klo di golang berubah jadi (2006-01-02 15:04:05)
	// ato bisa pake rfc bawaan ex: time.RFC3339
	formatter := "2006-01-02 15:04:05"
	parse, err := time.Parse(formatter, "2024-01-22 13:57:00")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(parse.AddDate(0, 1, 0))
	durationSecond := 100 * time.Second
	durationMiliSecond := 100 * time.Millisecond
	durationNano := durationSecond - durationMiliSecond
	fmt.Println(durationSecond)
	fmt.Println(durationMiliSecond)
	fmt.Println(durationNano)
	fmt.Printf("duration: %d\n", durationNano)

	// reflection
	// untuk membaca format metadata
	// dan bisa digunakan untuk simple validation juga dengan membaca struct tag
	user := User{Name: "tes", Age: 20}
	sampleType := reflect.TypeOf(user)
	fmt.Println("type name", sampleType.Name())
	isValid := true
	for i := 0; i < sampleType.NumField(); i++ {
		fmt.Println("field", sampleType.Field(i))
		fmt.Println("fieldName", sampleType.Field(i).Name)
		fmt.Println("fieldTag", sampleType.Field(i).Tag)
		if sampleType.Field(i).Tag.Get("required") == "true" {
			fmt.Println(reflect.ValueOf(user).Field(i).Interface())
			val := reflect.ValueOf(user).Field(i).Interface()
			if val == "" {
				isValid = false
			}
		}
	}
	if isValid {
		fmt.Println("user valid")
	}

	var patternRegexp = regexp.MustCompile(`e([a-z])o`)
	patternRegexp.MatchString("kharis")
	patternRegexp.MatchString("eko")
	patternRegexp.FindAllString("khai eko ego ake", 1)

	runEncode()
	csvReader()
	csvWriter()
	stringReader()
	createNewFile("./standard_lib/sample.log", "isian file sample")
	contentFile, _ := readFile("./standard_lib/sample.log")
	fmt.Println(contentFile)

	number := []int{1, 2, 3, 4, 5}
	days := []string{"senin", "selasa", "rabu"}
	fmt.Println(slices.Min(number))
	fmt.Println(slices.Max(number))
	fmt.Println(slices.Contains(days, "jumat"))
	fmt.Println(slices.Index(days, "jumat"))
	fmt.Println(slices.Index(days, "selasa"))

	fmt.Println(path.Dir("home/main.go"))
	fmt.Println(path.Base("home/ext/readme.md"))
	fmt.Println(path.Ext("./main.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))

	fmt.Println(filepath.Dir("home/main.go"))
	fmt.Println(filepath.Base("home/ext/readme.md"))
	fmt.Println(filepath.Ext("./main.go"))
	fmt.Println(filepath.IsAbs("home/ext/user/me.md"))
	fmt.Println(filepath.IsLocal("home/ext/user/me.md"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}

type User struct {
	Name string `required:"true" max:"10"` // struct tag
	Age  int    `required:"true"`
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func runErrors() {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println(err.Error())
		} else if errors.Is(err, NotFoundError) {
			fmt.Println(err.Error())
		} else {
			fmt.Println("unknown error")
		}
	}
}

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "1" {
		return NotFoundError
	}
	return nil
}

func runFlag() {
	host := flag.String("host", "localhost", "message")
	port := flag.Int("port", 3306, "db port")
	dbName := flag.String("db", "learn", "db name")

	flag.Parse()
	fmt.Println("host", *host)
	fmt.Println("port", *port)
	fmt.Println("dbName", *dbName)
}

func runEncode() {
	var encode = base64.StdEncoding.EncodeToString([]byte("Kharisma Nanda Wardhana"))
	fmt.Println(encode)
	fmt.Println(runDecode(encode))
}

func runDecode(encoded string) string {
	result, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err.Error()
	}
	return string(result)
}

func csvReader() {
	csvString := "aris,ara,\neri,aka,\narya,arin,\narga,anto"
	reader := csv.NewReader(strings.NewReader(csvString))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}

func csvWriter() {
	writer := csv.NewWriter(os.Stdout)
	err := writer.Write([]string{"TEST", "ARR"})
	if err != nil {
		fmt.Println(err.Error())
	}
	writer.Flush()
}

func stringReader() {
	input := strings.NewReader("ini adalah data\nyang akan dibaca")
	reader := bufio.NewReader(input)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	stringWriter()
}

func stringWriter() {
	w := bufio.NewWriter(os.Stdout)
	_, err := w.WriteString("ini Writer dari bufio\n")
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Flush()
}

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var content = ""
	for {
		line, _, err2 := reader.ReadLine()
		if err2 == io.EOF {
			break
		}
		content += string(line)
	}
	return content, nil
}
