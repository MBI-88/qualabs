package src

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type provider struct {
	ContenModule string `json:"content_module"`
	AuthModule   string `json:"auth_module"`
}

type user struct {
	Name     string   `json:"name"`
	Provider provider `json:"provider"`
	Id       string   `json:"id"`
}

type managerUsers struct {
	Users         []user
	AuthModule    map[string][]string
	ContentModule map[string][]string
}

// LoadData loads user files fron the path provided
//
// # Parameters
//
// path: valid path to the user files
//
// # Returns
//
// ture/false: if the datas were loaded return true, otherwise false
func (mg *managerUsers) LoadData(path string) bool {
	var (
		files = make([]string, 0, 25)
		usr   user
	)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, info.Name())
		}

		return nil
	})

	if err != nil {
		return false
	}

	for _, val := range files {
		if val != "" {
			p := fmt.Sprintf("%s/%s", path, val)
			f, err := os.Open(p)
			if err != nil {
				return false
			}
			body, err := io.ReadAll(f)
			if err != nil {
				return false
			}
			if err := json.Unmarshal(body, &usr); err == nil {
				usr.Id = fmt.Sprintf("./%s", val)
				mg.Users = append(mg.Users, usr)
			}
			f.Close()
		}

	}
	return true
}

// Solution A resolves the qualabs challenge A
//
// # Parameters
//
// # Returns
//
// (body, error): body (the solution required), error (if any error took place)
func (mg *managerUsers) SolutionA() ([]byte, error) {
	mg.feedMaps()
	response := make(map[string]map[string][]string)
	response["auth_module"] = mg.AuthModule
	response["content_module"] = mg.ContentModule

	body, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Solution B resolves the qualabs challenge B
//
// # Parameters
//
// # Returns
//
// minGroup: the smallest group formed by users of all modules
func (mg *managerUsers) SolutionB() []string {
	var (
		totalModule = make([]string, 0, 15)
		minGroup    = make([]string, len(mg.Users), 20)
		trackGroup  = make([]user, 0, 20)
		backtrack   func(index int)
	)

	mg.feedMaps()
	for key := range mg.AuthModule {
		totalModule = append(totalModule, key)
	}
	for key := range mg.ContentModule {
		totalModule = append(totalModule, key)
	}

	backtrack = func(index int) {
		if mg.checkTotalModules(trackGroup) == len(totalModule) && len(trackGroup) < len(minGroup) {
			minGroup = minGroup[:0]
			for _, us := range trackGroup {
				minGroup = append(minGroup, us.Id)
			}
			return
		}

		for i := index; i < len(mg.Users); i++ {
			if len(trackGroup) >= len(minGroup) {
				break
			}
			trackGroup = append(trackGroup, mg.Users[i])
			backtrack(i+1)
			trackGroup = trackGroup[:len(trackGroup)-1]

		}
	}

	backtrack(0)
	return minGroup
}

// feedMaps creates maps of two differents modules
//
// # Parameters
//
// # Returns
//
// void
func (mg *managerUsers) feedMaps() {
	mg.AuthModule = make(map[string][]string)
	mg.ContentModule = make(map[string][]string)
	for _, u := range mg.Users {
		mg.AuthModule[u.Provider.AuthModule] = append(mg.AuthModule[u.Provider.AuthModule], u.Id)
		mg.ContentModule[u.Provider.ContenModule] = append(mg.ContentModule[u.Provider.ContenModule], u.Id)
	}
}

// checkTotalModules checkes only uninq modules
//
// # Parameters
//
// us: user's array
//
// # Returns
//
// lenthg of the array of uniq modules
func (mg *managerUsers) checkTotalModules(us []user) int {
	totalModule := make([]string, 0, 20)
	for _, u := range us {
		if !mg.checkModuleInModules(u.Provider.AuthModule, totalModule) {
			totalModule = append(totalModule, u.Provider.AuthModule)
		}
		if !mg.checkModuleInModules(u.Provider.ContenModule, totalModule) {
			totalModule = append(totalModule, u.Provider.ContenModule)
		}
	}
	return len(totalModule)
}

// checkModuleInModules checkes module in moduels array
//
// # Parameters
//
// module: the module to check
// total: array of uninq modules
//
// # Returns
//
// (true/false): true (if module is in modules), false otherwise
func (mg *managerUsers) checkModuleInModules(module string, total []string) bool {
	for _, m := range total {
		if module == m {
			return true
		}
	}
	return false
}
