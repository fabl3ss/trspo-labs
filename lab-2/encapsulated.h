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
    double _value;;
public:
    EncapsulatedFloatWithMutex() : _value(0) {}
    const float getValue() override 
    { 
        _mtx.lock();
        auto val = this->_value;
        _mtx.unlock();
        return val; 
    }
    const void setValue(float newVal) override
    { 
        _mtx.lock();
        this->_value = newVal; 
        _mtx.unlock();
    } 
};

