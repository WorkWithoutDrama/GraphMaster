package main

//подключение к базе

func AddTask(id_project name_task) {
INSERT INTO Task (id_project, name_task) VALUES (@id_project, @name_task)}

func UpdateTask(id_project name_task param val) {
UPDATE Task @param = @val WHERE id_project = @id_project and name_task = @name_task}

func DeleteTask(id_project name_task) {
DELETE FROM Task WHERE id_project = @id_project and name_task = @name_task}

func AddSource(id_project name_source) {
INSERT INTO Source (id_project, name_source) VALUES (@id_project, @name_source)}

func UpdateSource(id_project name_source param val) {
UPDATE books Source @param = @val WHERE id_project = @id_project and name_task = @name_source}

func DeleteSource(id_project name_source) {
DELETE FROM Source WHERE id_project = @id_project and name_task = @name_source}

func GetInSourcesFromTask(id_project name_task) {
Select inSources FROM Task WHERE id_project = @id_project and name_task = @name_task}

func GetOutSourcesFromTask(id_project name_task) {
Select outSources FROM Task WHERE id_project = @id_project and name_task = @name_task}

func GetInTasksFromSource(id_project name_source) {
Select inTasks FROM Source WHERE id_project = @id_project and name_task = @name_source}

func GetOutTasksFromSource(id_project name_source) {
Select outTasks FROM Source WHERE id_project = @id_project and name_task = @name_sources}

func AddInSourcesFromTask(id_project name_task source) {
inSources := GetInSourcesFromTask(@id_project @name_task)
inSources = append(inSources, @source)
UpdateTask(@id_project @name_task inSources_task inSources)}

func AddOutSourcesFromTask(id_project name_task source) {
outSources := GetOutSourcesFromTask(@id_project @name_task)
outSources = append(outSources, @source)
UpdateTask(@id_project @name_task outSources_task outSources)}

func DeleteInSourcesFromTask(id_project name_task source) {
inSources := GetInSourcesFromTask(@id_project @name_task)
inSources = Remove(inSources, @source)
UpdateTask(@id_project @name_task inSources_task inSources)}

func DeleteOutSourcesFromTask(id_project name_task source) {
outSources := GetOutSourcesFromTask(@id_project @name_task)
outSources = Remove(outSources, @source)
UpdateTask(@id_project @name_task outSources_task outSources)}

func AddInTasksFromSource(id_project name_source task) {
inTasks := GetInTasksFromSource(@id_project @name_source)
inTasks = append(inTasks, @source)
UpdateTask(@id_project @name_source inTasks_source inSources)}

func AddOutTasksFromSource(id_project name_source task) {
outSources := GetOutTasksFromSource(@id_project @name_source)
outSources = append(outSources, @source)
UpdateTask(@id_project @name_source outTasks_source outSources)}

func DeleteInTasksFromSource(id_project name_source task) {
inTasks := GetInTasksFromSource(@id_project @name_source)
inTasks = Remove(inTasks, @source)
UpdateTask(@id_project @name_source inTasks_source inSources)}

func DeleteOutTasksFromSource(id_project name_source task) {
outSources := GetOutTasksFromSource(@id_project @name_source)
outSources = Remove(outSources, @source)
UpdateTask(@id_project @name_source outTasks_source outSources)}
