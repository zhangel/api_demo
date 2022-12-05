# import_db data import tools

### 1、数据导入工具说明

命令名称：		import_db

支持数据库类型： MySQL

支持参数说明：

--path         要导入的文件路径，目前仅支持XML格式文件，例如：--path=./file_path

--go_num  导入数据支持的协程数，例如：--go_num=10,  非必填参数，默认为1

--second    多久写入一条数据，例如：--second=10,每10秒写入一条数据，非必填参数，默认为0

**导入数据**

--path 支持文件路径例子：

./import_db --path=./655e3f39-3a52-41d9-bd21-21ea994b75aa.xml --go_num=10 --second=2

--path 支持包含xml文件的目录：

./import_db --path=./  --go_num=10 --second=2
