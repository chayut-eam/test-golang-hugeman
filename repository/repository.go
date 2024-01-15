package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	er "github.com/chayut-eam/test-golang-hugeman/error"
	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/chayut-eam/test-golang-hugeman/utils"
)

type Repository interface {
	GetAll() (*[]model.Data, error)
	Create(req *model.Data) error
	SearchByTilteOrDes(searchKey string) (*[]model.Data, error)
	UpdateById(id string, req *model.Data) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() (*[]model.Data, error) {

	filePath, err := getJsonFile()
	if err != nil {
		return nil, err
	}
	jsonData, err := ioutil.ReadFile(*filePath)
	if err != nil {
		return nil, err
	}

	if len(jsonData) <= 0 {
		b := []byte("[]")
		jsonData = b
	}

	data := []model.Data{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Title < data[j].Title || data[i].CreatedAt < data[j].CreatedAt || data[i].Status < data[j].Status
	})

	return &data, nil
}

func (r *repository) Create(req *model.Data) error {
	data, err := r.GetAll()
	if err != nil {
		return err
	}

	if binarySearchByTitle(*data, req.Title) {
		return errors.New(fmt.Sprintf("title '%s' is exist", req.Title))
	}

	newData := *data
	newData = append(newData, *req)
	if err = writeJsonData(newData); err != nil {
		return err
	}

	return nil
}

func getJsonFile() (*string, error) {
	currentPath, err := utils.GetRelativePath()
	if err != nil {
		return nil, err
	}
	filepath := filepath.Join(currentPath, "data", "data.json")
	return &filepath, nil
}

func writeJsonData(newData []model.Data) error {
	jsonData, err := json.MarshalIndent(newData, "", "    ")
	if err != nil {
		return err
	}

	filePath, err := getJsonFile()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(*filePath, os.O_RDWR, 7777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func indexDataById(data *[]model.Data) map[string]model.Data {
	hashMap := make(map[string]model.Data)
	for _, v := range *data {
		hashMap[v.ID] = v
	}

	return hashMap
}

func indexDataByTilteOrDes(data *[]model.Data) map[string][]model.Data {
	hashMap := make(map[string][]model.Data)
	for _, v := range *data {
		hashMap[v.Title] = append(hashMap[v.Title], v)
		hashMap[v.Description] = append(hashMap[v.Description], v)
	}

	return hashMap
}

func (r *repository) SearchByTilteOrDes(searchValue string) (*[]model.Data, error) {
	data, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	mapData := indexDataByTilteOrDes(data)

	if v, ok := mapData[searchValue]; ok {
		return &v, nil
	}

	return nil, er.NewDefinedError("Data Not Found", nil, "404")
}

func (r *repository) UpdateById(id string, req *model.Data) error {
	resultData := []model.Data{}
	data, err := r.GetAll()
	if err != nil {
		return err
	}

	mapData := indexDataById(data)

	if _, ok := mapData[id]; !ok {
		return er.NewDefinedError("Data Not Found", nil, "404")
	}
	mapData[id] = *req

	for _, v := range mapData {
		resultData = append(resultData, v)
	}

	if err := writeJsonData(resultData); err != nil {
		return err
	}

	return nil
}

func binarySearchByTitle(array []model.Data, to_search string) bool {
	found := false
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid].Title == to_search {
			found = true
			break
		}
		if array[mid].Title > to_search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return found
}
