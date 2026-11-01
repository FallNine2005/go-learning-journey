"使用linux系统成功的创建了目录和文件，学会了vim，touch，rm，mkdir等指令的操作。"
"成功使用cp指令更新了目录结构，修复了git库。"
"创建了mv1.go文件并使用mv指令改名为mv.go。"
"用cat指令查看README.md。"
"建立了big_file.txt并用less功能交互。"
"学会了head，tail，echo的使用。"
"使用tail -f实现了对bigfile的追踪，并用echo和>\>>实现了输出转移并追踪。"
"使用ls -l >实现了将文件目录覆盖bigfile文件，并学习了cat >，echo >>等指令。"
"cal指令显示日期并覆盖bigfile"
"运用ln -s指令建立了软连接，用test_txt指向了bigfile。"
"history"

|第二天的学习内容——接下来同步github会在一天的内容完结时一同发送|

“学会了使用date指令查看时间，cal指令查看年历。”
“学会使用find指令查找并将结果写入bigfile。”
“使用updatedb完成了遍历，并用locate查找”
“使用| gerp完成内容“March”的查找和显示行号查找Mo并写入bigfile”
“使用gzip对bigfile压缩并用gunzip解压”
“用zip -r将整个linux learning文件夹压缩为testzip.zip”
“使用tar -czvf将01，02，linuxlearning文件夹打包压缩为resttar.tar.zip，随后解压到新建testtar文件夹内”
“使用ls -ahl 查看文件的组和所有者，使用useradd -g A testA来建立了一个属于组A的测试账号testA，切换账号并生成了一个test.txt文件，查询到其属于组A，使用chgrp将其组别改为B。使用usermod将testa改为组B再使用chgrp指令将testa家目录下的所有文件改为组B。”

|第三天的学习内容|

“使用ll指令查询文件的权限并打印去bigfiles，学习了，-，l，d，c，b的对应表示内容，学会了0-9位前置字符代表的权限含义，使用chmod修改bigfile权限为rwxr-xr--。”
“学会了chown改变所有者和组的用法，将testa的文件夹所有者改为fallnine2005”
“学会了chgrp（先前已经使用过）”
“完成了警察土匪游戏实例操作————设置C，D组，设置用户police，bandit并分别加入C,D组，使用police创立文件bandit_will_die，设置权限640，随后修改为764，再将土匪bandit投降从D组移入C组检验是否可以查看警察police的文件。结果：成功。并写入bandit success to read,write this file。”
“创建了testcrontab.txt文件，并使用crontab指令每天调用一次文件。每天会在下午2点时自动写入read.me。
创建了testcrontab.sh脚本，并每天下午2点自动写入当前的时间日期进testcrontab2.txt文件中。”
“使用ps -ef | gerp atd来检测atd进程是否在运行，写入一个测试用at命令，将会在5点写入一个一次性测试语句。”Sun Nov  1 17:00:00 CST 2026
succeed to use at

|第四天的内容|