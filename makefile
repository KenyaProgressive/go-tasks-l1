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

clean-all:
	cd 1 && rm -f task1.out && cd ..
	cd 2 && rm -f task2.out && cd ..
	cd 3 && rm -f task3.out && cd ..
	cd 4 && rm -f task4.out && cd ..
	cd 5 && rm -f task5.out && cd ..
	cd 6 && rm -f task6.out && cd ..