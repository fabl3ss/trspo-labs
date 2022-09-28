#include <future>
#include <iomanip>
#include <iostream>

#include "application.h"

std::uniform_int_distribution<unsigned int> randomThreadsAmount (
	MIN_THREADS_AMOUNT, 
	MAX_THREADS_AMOUNT	
);
std::uniform_int_distribution<unsigned int> randomIterationsAmount(
	MIN_ITERATIONS_AMOUNT, 
	MAX_ITERATIONS_AMOUNT	
);
std::uniform_real_distribution<float> randomValue;

Application::Application() : _mt(std::mt19937(this->_rd())) 
{
	this->a = new EncapsulatedFloatWithMutex;
	this->b = new EncapsulatedFloatWithMutex;

    this->_workersAmount = randomThreadsAmount(this->_mt);
    this->_firstWorkIterations = randomIterationsAmount(this->_mt);
    this->_secondWorkIterations = randomIterationsAmount(this->_mt);
}

void Application::createWorkers() 
{
    auto firstJobWorkersAmount = this->_workersAmount / 2;
    auto secondJobWorkersAmount = this->_workersAmount - firstJobWorkersAmount;
	this->createFirstJobThreads(firstJobWorkersAmount);
	this->createSecondJobThreads(secondJobWorkersAmount);
}

void Application::waitAllWorkers() 
{
    for (auto& a : this->_workers) {
        a.join();
    }
}

void Application::Run()
{
    createWorkers();
    waitAllWorkers();
    std::cout << "a._value==[" << this->a->getValue() << "]";
    std::cout << "a._value==[" << this->b->getValue() << "]";
}

void Application::createFirstJobThreads(unsigned int threadsAmount) 
{
    auto work = std::bind(
        &Application::firstJob, 
        std::placeholders::_1,
        std::placeholders::_2
    );

    for (size_t i = 0; i < threadsAmount; ++i) {
        this->_workers.emplace_back(
            work, 
            this,
            this->_firstWorkIterations
        );
    }
}

void Application::createSecondJobThreads(unsigned int threadsAmount) 
{
    auto work = std::bind(
        &Application::secondWork, 
        std::placeholders::_1,
        std::placeholders::_2
    );

    for (size_t i = 0; i < threadsAmount; ++i) {
        this->_workers.emplace_back(
            work, 
            this, 
            this->_secondWorkIterations 
        );
    }
}

void Application::firstJob(size_t iters)
{
    for (size_t i = 0; i < iters; i++) {
        this->printWorkStatus(__FUNCTION__);
	    a->setValue(a->getValue() + randomValue(this->_mt));
	    b->setValue(b->getValue() + randomValue(this->_mt));
    }
}

void Application::secondWork(size_t iters) 
{
    for (size_t i = 0; i < iters; i++) {
        this->printWorkStatus(__FUNCTION__);
	    b->setValue(b->getValue() + 1);
	    a->setValue(a->getValue() + 1);
    }
}

void Application::printWorkStatus(std::string workName)
{
    this->_mtx.lock();
    auto threadIdMaxLenght = 6;
    std::cout << workName 
        << "\tThreadId: "<< std::setw(threadIdMaxLenght) << std::this_thread::get_id()
        << "\ta._value==[" << a->getValue() << "]" 
        << "\tb._value==[" << b->getValue() << "]"
        <<  std::endl;
    this->_mtx.unlock();
}

