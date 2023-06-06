package main

import "fmt"

func StartSpecTask(id IdTask) {
	task := getSpecTask(id.Id)
	shipment := getSpecShipment(int(task.Shipment_id))
	card := getSpecCard(int(task.Card_id))
	taskStart := ReadyTask{
		Task: task,
		Shipment: shipment,
		Card: card,
	}
	fmt.Println(taskStart)
	 if taskStart.Task.Shop == "Palace" {
	 	palace(taskStart)
	 }else{
	 	supreme(taskStart)
	 }
}
