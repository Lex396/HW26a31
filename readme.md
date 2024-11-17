1.
	git init
2.
	git add .
	git commit -m "beginning"
3.
	git tag -a V0.1 -m "start"
4.
	git remote add origin https://github.com/Lex396/HW26a31.git
	git branch -M main
	git push origin main
	git push V0.1
5. 
	git checkout -b develop main
 	git push origin develop
 6.
 	git checkout -b feature/add_log develop
 	git add main.go
 7.
 	git commit -m "added logging"
 	git push origin feature/add_log
 	git checkout develop
 	git merge --no-ff feature/add_log
 	git push origin develop
 8.
 	git checkout -b release develop
 	git tag -a V1.0 -m "V1.0"
 9.
 	git push origin release
 	git push origin V1.0
 	git checkout main
 	git merge --no-ff release
 	git push origin main
10.
	Dockerfile, описывающий двухступенчатую сборку образа
11.
	git add main.go dockerfile readme.md
	git commit -m "add logingadded logging to the console and to the file"
	git checkout main
	git merge --no-ff develop
	git tag -a V1.1 -m "V1.1"
12.
	git push origin develop
	git push origin main
	git push origin V1.1