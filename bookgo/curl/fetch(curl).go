
// Программа которая выводит  HTML данные веб сайт, похожа на CURL 
// В случае ошибки HTTP-запроса fetch сообщает о том, что произошло:
//  $./fetch http://bad.gopl.io
/// fetch: Get http://bad.gopl.io: dial tcp: lookup bad.gopl.io: no such host
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func curl(){
	
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch:%v\n" , err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch: чтение %s:%v\n", url, err)
			os.Exit(1)
		}
		
		fmt.Printf("%s", b)
		fmt.Printf(resp.Request.Proto,"status",resp.StatusCode)
	}
}

///Упражнение 1.7. Вызов функции io. Copy (dst,src) выполняет чтение src  и
///запись в dst. Воспользуйтесь ею вместо ioutil. ReadAll для копирования тела
///ответа в поток os. Stdout без необходимости выделения достаточно большого для
///хранения всего ответа буфера. Не забудьте проверить, не произошла ли ошибка при
///вызове io. Сору.
///Упражнение 1.8. Измените программу fetch так, чтобы к каждому аргументу
///URL автоматически добавлялся префикс http:// в случае отсутствия в нем таково­го
/// Можете воспользоваться функцией strings. HasPrefix.
///упражнение 1.9. Измените программу fetch так, чтобы она выводила код состо­-
/// -яния HTTP, содержащийся в resp.Status .