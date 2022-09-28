#include <mutex>
#include <random>
#include <thread>
#include <functional>

#include "encapsulated.h"

#define MIN_THREADS_AMOUNT 10
#define MAX_THREADS_AMOUNT 20
#define MIN_ITERATIONS_AMOUNT 10000
#define MAX_ITERATIONS_AMOUNT 20000

class Application
{
	std::mutex _mtx;
	std::mt19937 _mt;
	std::random_device _rd;

	std::vector<std::thread> _workers;

	unsigned int _workersAmount;
	unsigned int _firstWorkIterations;
	unsigned int _secondWorkIterations;

	EncapsulatedFloatWithMutex* a;
	EncapsulatedFloatWithMutex* b;

	void firstJob(size_t iters);
	void secondWork(size_t iters);
	void createFirstJobThreads(unsigned int amount);
	void createSecondJobThreads(unsigned int threadsAmount);
	void printWorkStatus(std::string workName);
	void createWorkers();
	void waitAllWorkers();
public:
	Application();
	void Run();
};

