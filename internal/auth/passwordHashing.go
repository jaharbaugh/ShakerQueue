package auth

import(

)

func CheckPasswordHash(password string, hash string) (bool, error){
	
	return argon2id.ComparePasswordAndHash(password, hash)

}