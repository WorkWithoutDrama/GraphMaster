package main

func AddTaskFromProject(id_project name_task) {
AddTask (@id_project, @name_task)}
//отправка сообщения о добавлении AddTaskSoT (@id_project, @name_task)

func AddSourceFromProject(id_project source_task) {
AddSource (@id_project, @name_source)}
//отправка сообщения о добавлении AddSourceSot (@id_project, @name_source)

func DeleteTaskFromProject(id_project name_task) {
inSources := GetInSourcesFromTask(@id_project @name_task)
for i, source := range inSources{
   go DeleteOutTasksFromSource(id_project source name_task)}
//отправка сообщения об удалении OutTaskFromSource
outSources := GetOutSourcesFromTask(@id_project @name_task)
for i, source := range outSources{
   go DeleteInTasksFromSource(id_project source name_task)}
//отправка сообщения об удалении InTaskFromSource
DeleteTask(id_project name_task)}

func DeleteSourcesFromProject(id_project name_source) {
inTasks := GetInTasksFromSource(@id_project @name_source)
for i, task := range inTasks{
   go DeleteOutSourcesFromTask(id_project task name_source)}
//отправка сообщения об удалении OutSourcesFromTask
outSources := GetOutSourcesFromTask(@id_project @name_source)
for i, task := range outTasks{
//если статус таски ожидает присвоить доступен и отправка сообщения об изменении Task//отправка сообщения об удалении
   go DeleteInSourcesFromTask(id_project source name_source)}
//отправка сообщения об удалении InSourcesFromTask
DeleteSource(id_project name_source)}

func AddInSourcesFromProject(id_project task source) {
AddInSourcesFromTask(id_project task source)
//отправка сообщения об изменении InSourcesFromTask
//если есть
AddOutTasksFromSource(id_project source task)
//отправка сообщения об изменении OutTasksFromSource
}

func AddOutSourcesFromProject(id_project task source) {
AddOutSourcesFromTask(id_project task source)
//отправка сообщения об изменении OutSourcesFromTask
//если есть
AddInTasksFromSource(id_project source task)
//отправка сообщения об изменении InTasksFromSource
}

func DeleteInSourcesFromProject(id_project task source) {
DeleteInSourcesFromTask(id_project task source)
//проверка статуса
//отправка сообщения об изменении InSourcesFromTask
DeleteOutTasksFromSource(id_project source task)
//отправка сообщения об изменении OutTasksFromSource
}

func DeleteOutSourcesFromProject(id_project task source) {
DeleteOutSourcesFromTask(id_project task source)
//отправка сообщения об изменении OutSourcesFromTask
DeleteInTasksFromSource(id_project source task)
//отправка сообщения об изменении InTasksFromSource
}

func AddInTasksSourcesProject(id_project source task) {
//если есть
AddInTasksFromSource(id_project source task)
//отправка сообщения об изменении InTasksFromSource
AddOutSourcesFromTask(id_project task source)
//отправка сообщения об изменении OutSourcesFromTask
}

func AddOutSourcesFromProject(id_project source task) {
//если есть
AddOutTasksFromSource(id_project source task)
//отправка сообщения об изменении OutTasksFromSource
AddInSourcesFromTask(id_project task source)
//отправка сообщения об изменении InSourcesFromTask
}

func DeleteInSourcesFromProject(id_project source task) {
DeleteInTasksFromSource(id_project source task)
//отправка сообщения об изменении InTasksFromSource
DeleteOutSourcesFromTask(id_project task source)
//отправка сообщения об изменении OutSourcesFromTask
}

func DeleteOutSourcesFromProject(id_project source task) {
DeleteOutTasksFromSource(id_project source task)
//отправка сообщения об изменении OutTasksFromSource
DeleteInSourcesFromTask(id_project task source)
//проверка статуса
//отправка сообщения об изменении InSourcesFromTask
}

//копирование
//name autor

//копирование
//name autor
