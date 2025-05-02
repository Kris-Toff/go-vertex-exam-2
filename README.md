# Go Vertex Exam 2 #

### Code Exam Intructions: ###

- Create an API that takes 1 input number
- Then get the next seven prime numbers
- and adds those to the original number
- Then output total
- Do this 10 times numbers after the input number

#### How to setup locally ####
- git pull `git@github.com:Kris-Toff/go-vertex-exam-2.git`
- open terminal and go to directory
- run `go run cmd/go_vertex_exam_2/main.go` this will start the local server 

#### Postman setup ####
- Method: `POST`
- URL: http://localhost:8080/
- In Body tab, select `x-www-form-urlencoded`
- In the form add Key: `number` and value, the value should be any positive whole number 
- Clicking send will return the array of objects containing `number`, `prime_numbers`, and `total`

<br><br>
Screenshot: 
![Postman Example](/screenshots/go_vertex_exam_2.jpg)