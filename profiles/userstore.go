// pipebomb/profiles/userstore.go
package profiles

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"os"
)

func init() {
	if _, err := os.Stat("users.json"); os.IsNotExist(err) {
		file, err := os.Create("users.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		emptyUsers := []User{}
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", " ")
		err = encoder.Encode(emptyUsers)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	file, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Loaded %d users from users.json\n", len(Users))

	// Update number of loaded JSON items when a user is added or deleted
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "DELETE" {
			file, err := os.Open("users.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&Users)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Loaded %d users from users.json\n", len(Users))
		}
	})
}

// saveUsers saves the current users list to the users.json file
func SaveUsers() {
	dataFile, err := os.Create("users.json")
	if err != nil {
		color.Red(err.Error())
		return
	}
	defer dataFile.Close()
	encoder := json.NewEncoder(dataFile)
	encoder.SetIndent("", " ")
	err = encoder.Encode(Users)
	if err != nil {
		color.Red(err.Error())
		return
	}
	color.Green("Users saved successfully")
}

// FindUserByUsername finds a user by their username in the users list
func FindUserByUsername(username string) (*User, error) {
	for _, user := range Users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

// AddUser adds a new user to the users list
func AddUser(user User) {
	Users = append(Users, user)
	SaveUsers()
}

// UpdateUser updates a user's data in the users list
func UpdateUser(updatedUser User) error {
	for i, user := range Users {
		if user.Username == updatedUser.Username {
			Users[i] = updatedUser
			SaveUsers()
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

// DeleteUser deletes a user from the users list
func DeleteUser(username string) (User, error) {
	if username != "Space Mommy" {
		for i, user := range Users {
			if user.Username == username {
				Users = append(Users[:i], Users[i+1:]...)
				SaveUsers()
				return user, nil
			}
		}
	}
	return User{}, errors.New("user not found")
}

func MergeUpdated(user *User, updatedUser User) {
	if updatedUser.Username != "" {
		user.Username = updatedUser.Username
	}
	if updatedUser.Avatar != "" {
		user.Avatar = updatedUser.Avatar
	}
	profile := &user.Profile
	updatedProfile := &updatedUser.Profile
	if updatedProfile.Name != "" {
		profile.Name = updatedProfile.Name
	}
	if updatedProfile.Image != "" {
		profile.Image = updatedProfile.Image
	}
	if updatedProfile.Bio != "" {
		profile.Bio = updatedProfile.Bio
	}
	if updatedProfile.Philosophy != "" {
		profile.Philosophy = updatedProfile.Philosophy
	}
}
