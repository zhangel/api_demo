# convert  tools

### 1、XML 、json format file to csv

命令名称：		convert

原始文件类型： XML、JSON

支持参数说明：

--src           源文件目录，里边可以放包含xml、json的文件，参数示例：--src=./src_path

--dest        目的文件目录，转换成CSV文件的目标目录，如果没有会创建，参数示例：--dest=./dest_path

--file_type 支持的文件类型，目前支持XML、JSON两种，参数示例： --file_type=xml,json 

--go_num  导入数据支持的协程数，例如：--go_num=10,  非必填参数，默认为1

--second    多久写入一条数据，例如：--second=10,每10秒写入一条数据，非必填参数，默认为0

**转换示例**

./convert --src=./ --dest=./aaa --file_type=xml,json

以上命令会将./目录下的xml、json格式文件转成csv文件，转存到aaa目录下，如果该目录不存在会创建。
