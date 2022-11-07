package core

import "testing"

func TestConvertInt(t *testing.T) {
	ret, err := Int(-100)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. -100 -> %d", ret)
	}
	ret, err = Int(false)
	if err != nil {
		t.Logf("convert int failed:%s", err)
	} else {
		t.Errorf("convert int success.false -> %d", ret)
	}
	ret, err = Int(uint(100))
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. uint(100) -> %d", ret)
	}
	ret, err = Int(100.1)
	if err != nil {
		t.Logf("convert int failed:%s", err)
	} else {
		t.Errorf("convert int success. 100.1 -> %d", ret)
	}
	ret, err = Int("100")
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. \"100\" -> %d", ret)
	}
	ret, err = Int([]byte("-100"))
	if err != nil {
		t.Errorf("convert int failed: %s", err)
	} else {
		t.Logf("convert int success. []byte(\"100\") -> %d", ret)
	}
}

func TestConvertUint(t *testing.T) {
	ret, err := Uint(-100)
	if err != nil {
		t.Logf("convert int failed:%s", err)
	} else {
		t.Errorf("convert int success. -100 -> %d", ret)
	}
	ret, err = Uint(false)
	if err != nil {
		t.Logf("convert int failed:%s", err)
	} else {
		t.Errorf("convert int success.false -> %d", ret)
	}
	ret, err = Uint(uint(100))
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. uint(100) -> %d", ret)
	}
	ret, err = Uint(100.1)
	if err != nil {
		t.Logf("convert int failed:%s", err)
	} else {
		t.Errorf("convert int success. 100.1 -> %d", ret)
	}
	ret, err = Uint("100")
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. \"100\" -> %d", ret)
	}
	ret, err = Uint([]byte("100"))
	if err != nil {
		t.Errorf("convert int failed: %s", err)
	} else {
		t.Logf("convert int success. []byte(\"100\") -> %d", ret)
	}
}

func TestConvertBool(t *testing.T) {
	ret, err := Bool(-100)
	if err != nil {
		t.Logf("convert int failed:%s", err)
	} else {
		t.Errorf("convert int success. -100 -> %t", ret)
	}
	ret, err = Bool(false)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success.false -> %t", ret)
	}
	ret, err = Bool(uint(100))
	if err != nil {
		t.Logf("convert int failed:%s", err)
	} else {
		t.Errorf("convert int success. uint(100) -> %t", ret)
	}
	ret, err = Bool(100.1)
	if err != nil {
		t.Logf("convert int failed:%s", err)
	} else {
		t.Errorf("convert int success. 100.1 -> %t", ret)
	}
	ret, err = Bool("true")
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. \"true\" -> %t", ret)
	}
	ret, err = Bool([]byte("false"))
	if err != nil {
		t.Errorf("convert int failed: %s", err)
	} else {
		t.Logf("convert int success. []byte(\"false\") -> %t", ret)
	}
}

func TestConvertFloat(t *testing.T) {
	ret, err := Float(-100)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. -100 -> %f", ret)
	}
	ret, err = Float(false)
	if err != nil {
		t.Logf("convert int failed:%s", err)
	} else {
		t.Errorf("convert int success.false -> %f", ret)
	}
	ret, err = Float(uint(100))
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. uint(100) -> %f", ret)
	}
	ret, err = Float(100.1)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. 100.1 -> %f", ret)
	}
	ret, err = Float("100.1")
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. \"100.1\" -> %f", ret)
	}
	ret, err = Float([]byte("100"))
	if err != nil {
		t.Errorf("convert int failed: %s", err)
	} else {
		t.Logf("convert int success. []byte(\"100\") -> %f", ret)
	}
}

func TestConvertString(t *testing.T) {
	ret, err := String(-100)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. -100 -> %s", ret)
	}
	ret, err = String(false)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success.false -> %s", ret)
	}
	ret, err = String(uint(100))
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. uint(100) -> %s", ret)
	}
	ret, err = String(100.1)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. 100.1 -> %s", ret)
	}
	ret, err = String("100.1")
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. \"100.1\" -> %s", ret)
	}
	ret, err = String([]byte("100"))
	if err != nil {
		t.Errorf("convert int failed: %s", err)
	} else {
		t.Logf("convert int success. []byte(\"100\") -> %s", ret)
	}
}

func TestConvertStringQuote(t *testing.T) {
	ret, err := StringQuote(-100)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. -100 -> %s", ret)
	}
	ret, err = StringQuote(false)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success.false -> %s", ret)
	}
	ret, err = StringQuote(uint(100))
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. uint(100) -> %s", ret)
	}
	ret, err = StringQuote(100.1)
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. 100.1 -> %s", ret)
	}
	ret, err = StringQuote("100.1")
	if err != nil {
		t.Errorf("convert int failed:%s", err)
	} else {
		t.Logf("convert int success. \"100.1\" -> %s", ret)
	}
	ret, err = StringQuote([]byte("100"))
	if err != nil {
		t.Errorf("convert int failed: %s", err)
	} else {
		t.Logf("convert int success. []byte(\"100\") -> %s", ret)
	}
}
