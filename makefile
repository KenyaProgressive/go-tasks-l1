task1:
	cd 1 && go build -o task1.out . && ./task1.out

task2:
	cd 2 && go build -o task2.out . && ./task2.out

task3:
	cd 3 && go build -o task3.out . && ./task3.out && cd ..

task4:
	cd 4 && go build -o task4.out . && ./task4.out && cd ..

task5:
	cd 5 && go build -o task5.out . && ./task5.out && cd ..

task6:
	cd 6 && go build -o task6.out . && ./task6.out && cd ..

task7:
	cd 7 && go build -race -o task7.out . && ./task7.out && cd ..

task8:
	cd 8 && go build -o task8.out . && ./task8.out && cd ..

task9:
	cd 9 && go build -o task9.out . && ./task9.out && cd ..
	
task10:
	cd 10 && go build -o task10.out . && ./task10.out && cd ..

task11:
	cd 11 && go build -o task11.out . && ./task11.out && cd ..

task12:
	cd 12 && go build -o task12.out . && ./task12.out && cd ..

task13:
	cd 13 && go build -o task13.out . && ./task13.out && cd ..	

task14:
	cd 14 && go build -o task14.out . && ./task14.out && cd ..

task15:
	cd 15 && go build -o task15.out . && ./task15.out && cd ..

task16:
	cd 16 && go build -o task16.out . && ./task16.out && cd ..

clean-all:
	cd 1 && rm -f task1.out && cd ..
	cd 2 && rm -f task2.out && cd ..
	cd 3 && rm -f task3.out && cd ..
	cd 4 && rm -f task4.out && cd ..
	cd 5 && rm -f task5.out && cd ..
	cd 6 && rm -f task6.out && cd ..
	cd 7 && rm -f task7.out && cd ..
	cd 8 && rm -f task8.out && cd ..
	cd 9 && rm -f task9.out && cd ..
	cd 10 && rm -f task10.out && cd ..
	cd 11 && rm -f task11.out && cd ..
	cd 12 && rm -f task12.out && cd ..
	cd 13 && rm -f task13.out && cd ..
	cd 14 && rm -f task14.out && cd ..
	cd 15 && rm -f task15.out && cd ..
	cd 16 && rm -f task16.out && cd ..