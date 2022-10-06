template<typename _valueType>
class Encapsulated 
{
public:
    virtual const _valueType getValue() = 0;
    virtual const void setValue(_valueType newVal) = 0;
};

class EncapsulatedFloatWithMutex : public Encapsulated<float> 
{
private:
    std::mutex _mtx;
    float _value;
public:
    EncapsulatedFloatWithMutex() : _value(0) {}
    const float getValue() override 
    { 
        double val;
        _mtx.lock();
        try {
            val = this->_value;
        }
        catch (std::exception &e) {
            _mtx.unlock();
            return this->_value;
        };
        _mtx.unlock();
        return val; 
    }
    const void setValue(float newVal) override
    { 
        _mtx.lock();
        try {
            this->_value = newVal; 
        }
        catch (std::exception& e) {
            _mtx.unlock();
            return;
        }
        _mtx.unlock();
    } 
};

