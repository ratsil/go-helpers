package helpers

import (
	"fmt"
	"math"
	sc "strconv"
	"strings"
	t "time"

	. "github.com/ratsil/go-helpers/dbc/types"
)

// ToPDT .
func ToPDT(i interface{}) *t.Time {
	if nil == i {
		return nil
	}
	if p, b := i.(*t.Time); b {
		return p
	}
	if dt, b := i.(t.Time); b {
		return &dt
	}
	if sValue, b := i.(string); b {
		n := len(sValue)
		if 18 < n {
			if dt, err := t.Parse("2006-01-02T15:04:05-0700", sValue); nil == err {
				return &dt
			}
			if dt, err := t.Parse("2006-01-02 15:04:05.999999-07:00", sValue); nil == err {
				return &dt
			}
			if dt, err := t.Parse("2006-01-02 15:04:05.999999-07", sValue); nil == err {
				return &dt
			}
			if dt, err := t.Parse("2006-01-02 15:04:05.999-07", sValue); nil == err {
				return &dt
			}
			if dt, err := t.Parse("02.01.2006 15:04:05.999999-07:00", sValue); nil == err {
				return &dt
			}
			if dt, err := t.Parse("02.01.2006 15:04:05.999-07", sValue); nil == err {
				return &dt
			}
			if dt, err := t.Parse("02.01.2006 15:04:05", sValue); nil == err {
				return &dt
			}
		} else if 9 < n {
			if dt, err := t.Parse("02.01.2006", sValue); nil == err {
				return &dt
			}
			if dt, err := t.Parse("2006-01-02", sValue); nil == err {
				return &dt
			}
		} else {
			if dt, err := t.Parse("20060102", sValue); nil == err {
				return &dt
			}
		}
	}
	return nil
}

// ToDT .
func ToDT(i interface{}) t.Time {
	p := ToPDT(i)
	if nil == p {
		return DTNull
	}
	return *p
}

// ToPStr .
func ToPStr(i interface{}, a ...interface{}) *string {
	defer func() { recover() }()
	if nil == i {
		return nil
	}
	if s, b := i.(string); b {
		return &s
	}
	if p, b := i.(*string); b {
		return p
	}
	if a, b := i.([]byte); b {
		return &[]string{string(a)}[0]
	}
	if dt, b := i.(*t.Time); b {
		if 1 > len(a) {
			a = append(a, "2006-01-02 15:04:05.99")
		}
		s := dt.Format(a[0].(string))
		return &s
	}
	if dt, b := i.(t.Time); b {
		if 1 > len(a) {
			a = append(a, "2006-01-02 15:04:05.99")
		}
		s := dt.Format(a[0].(string))
		return &s
	}
	if n, b := i.(int8); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatInt(int64(n), 10)
		}
		return &s
	}
	if n, b := i.(int16); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatInt(int64(n), 10)
		}
		return &s
	}
	if n, b := i.(int); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatInt(int64(n), 10)
		}
		return &s
	}
	if n, b := i.(int64); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatInt(n, 10)
		}
		return &s
	}
	if n, b := i.(uint8); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatUint(uint64(n), 10)
		}
		return &s
	}
	if n, b := i.(uint16); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatUint(uint64(n), 10)
		}
		return &s
	}
	if n, b := i.(uint); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatUint(uint64(n), 10)
		}
		return &s
	}
	if n, b := i.(uint64); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatUint(n, 10)
		}
		return &s
	}
	if n, b := i.(ID); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatInt(int64(n), 10)
		}
		return &s
	}
	if n, b := i.(float64); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatFloat(n, 'f', -1, 64)
		}
		return &s
	}
	if bb, b := i.(bool); b {
		s := sc.FormatBool(bb)
		return &s
	}
	return nil
}

// ToStr .
func ToStr(i interface{}, a ...interface{}) string {
	p := ToPStr(i, a...)
	if nil == p {
		return ""
	}
	return *p
}

// ToInt64 .
func ToInt64(i interface{}) int64 {
	defer func() { recover() }()
	switch v := i.(type) {
	case nil:
		return math.MaxInt64
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		if n, err := sc.ParseInt(v, 10, 64); nil == err {
			return n
		}
	}
	return math.MaxInt64
}

// ToLong .
func ToLong(i interface{}) int64 {
	return ToInt64(i)
}

// ToInt16 .
func ToInt16(i interface{}) int16 {
	v := ToInt64(i)
	if v != math.MaxInt64 {
		return int16(v)
	}
	return math.MaxInt16
}

// ToInt8 .
func ToInt8(i interface{}) int16 {
	v := ToInt64(i)
	if v != math.MaxInt64 {
		return int16(v)
	}
	return math.MaxInt16
}

// ToShort .
func ToShort(i interface{}) int16 {
	return ToInt16(i)
}

// ToInt .
func ToInt(i interface{}) int {
	v := ToInt64(i)
	if v != math.MaxInt64 {
		return int(v)
	}
	return math.MaxInt32
}

// ToUInt64 .
func ToUInt64(i interface{}) uint64 {
	defer func() { recover() }()
	switch v := i.(type) {
	case nil:
		return math.MaxUint64
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint64:
		return v
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case string:
		if n, err := sc.ParseUint(v, 10, 64); nil == err {
			return n
		}
	}
	return math.MaxUint64
}

// ToULong .
func ToULong(i interface{}) uint64 {
	return ToUInt64(i)
}

// ToUInt32 .
func ToUInt32(i interface{}) uint32 {
	v := ToUInt64(i)
	if v != math.MaxInt64 {
		return uint32(v)
	}
	return math.MaxUint32
}

// ToUInt .
func ToUInt(i interface{}) uint {
	v := ToUInt64(i)
	if v != math.MaxInt64 {
		return uint(v)
	}
	return math.MaxUint
}

// ToUInt8 .
func ToUInt8(i interface{}) uint8 {
	v := ToUInt64(i)
	if v != math.MaxInt64 {
		return uint8(v)
	}
	return math.MaxUint8
}

// ToByte .
func ToByte(i interface{}) byte {
	return byte(ToUInt8(i))
}

// ToFloat64 .
func ToFloat64(i interface{}, a ...interface{}) (ret float64) {
	defer func() { recover() }()
	ret = math.MaxFloat64
	switch v := i.(type) {
	case int:
		ret = float64(v)
	case int8:
		ret = float64(v)
	case int16:
		ret = float64(v)
	case int64:
		ret = float64(v)
	case uint:
		ret = float64(v)
	case uint8:
		ret = float64(v)
	case uint16:
		ret = float64(v)
	case uint64:
		ret = float64(v)
	case float32:
		ret = float64(v)
	case float64:
		ret = v
	case string:
		if n, err := sc.ParseFloat(v, 64); nil == err {
			ret = n
		}
	}
	if math.MaxFloat64 > ret {
		if 0 < len(a) {
			if n, b := a[0].(int); b {
				n *= 10
				ret = math.Floor(ret*float64(n)) / float64(n)
			}
		}
	}
	return
}

// ToDouble .
func ToDouble(i interface{}, a ...interface{}) float64 {
	return ToFloat64(i, a...)
}

// ToID .
func ToID(i interface{}) ID {
	return ID(ToInt64(i))
}

// Round .
func Round(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}

// ToBool .
func ToBool(i interface{}) bool {
	if nil != i {
		if s, b := i.(string); b {
			s = strings.ToLower(strings.TrimSpace(s))
			return ("t" == s || "true" == s)
		}
	}
	return false
}

// ToPBool .
func ToPBool(i interface{}) *bool {
	if nil == i {
		return nil
	}
	if p, b := i.(*bool); b {
		return p
	}
	if bValue, b := i.(bool); b {
		return &bValue
	}
	return nil
}

/*
func ToByte(i interface{}) byte
{
	if i is byte)
		return (byte)i;
	byte nRetVal = byte.MaxValue;
	if nil != i)
	{
		try
		{
			nRetVal = Convert.ToByte(i);
		}
		catch { }
	}
	return nRetVal;
}
func short ToInt16(i interface{})
{
	if i is short)
		return (short)i;
	short nRetVal = short.MaxValue;
	if nil != i)
	{
		try
		{
			nRetVal = Convert.ToInt16(i);
		}
		catch { }
	}
	return nRetVal;
}
func ushort ToUInt16(i interface{})
{
	if i is ushort)
		return (ushort)i;
	ushort nRetVal = ushort.MaxValue;
	if nil != i)
	{
		try
		{
			nRetVal = Convert.ToUInt16(i);
		}
		catch { }
	}
	return nRetVal;
}
func short ToShort(i interface{})
{
	return i.ToInt16();
}
func ushort ToUShort(i interface{})
{
	return i.ToUInt16();
}
func int ToInt32(i interface{})
{
	if i is int)
		return (int)i;
	int nRetVal = int.MaxValue;
	if nil != i)
	{
		try
		{
			nRetVal = Convert.ToInt32(i);
		}
		catch { }
	}
	return nRetVal;
}
func uint ToUInt32(i interface{})
{
	if i is uint)
		return (uint)i;
	uint nRetVal = uint.MaxValue;
	if nil != i)
	{
		try
		{
			nRetVal = Convert.ToUInt32(i);
		}
		catch
		{
			if i is string && ((string)i).StartsWith("0x"))
			{
				try
				{
					nRetVal = Convert.ToUInt32((string)i, 16);
				}
				catch { }
			}
		}
	}
	return nRetVal;
}
func int ToInt(i interface{})
{
	return i.ToInt32();
}
func uint ToUInt(i interface{})
{
	return i.ToUInt32();
}
func long ToInt64(i interface{})
{
	if i is long)
		return (long)i;
	long nRetVal = long.MaxValue;
	if nil != i)
	{
		try
		{
			nRetVal = Convert.ToInt64(i);
		}
		catch { }
	}
	return nRetVal;
}
func ulong ToUInt64(i interface{})
{
	if i is ulong)
		return (ulong)i;
	ulong nRetVal = ulong.MaxValue;
	if nil != i)
	{
		try
		{
			nRetVal = Convert.ToUInt64(i);
		}
		catch { }
	}
	return nRetVal;
}
func long ToLong(i interface{})
{
	return i.ToInt64();
}
func ulong ToULong(i interface{})
{
	return i.ToUInt64();
}
func double ToDouble(this object i)
{
	return i.ToDouble(nil);
}
func double ToDouble(this object i, ushort nDecimals)
{
	return Math.Round(i.ToDouble(nil), nDecimals);
}
func double ToDouble(this object i, IFormatProvider iFormatProvider)
{
	if i is double)
		return (double)i;
	double nRetVal = double.MaxValue;
	if nil != i)
	{
		try
		{
			if(i is byte[])
				nRetVal = ((byte[])i).ToDouble(0, false);
			else
				nRetVal = Convert.ToDouble(i, iFormatProvider);
		}
		catch { }
	}
	return nRetVal;
}
func float ToFloat(i interface{})
{
	if i is float)
		return (float)i;
	float nRetVal = float.MaxValue;
	if nil != i)
	{
		try
		{
			if i is string)
			{
				string sValue = (string)i;
				sValue = sValue.Replace(".", ",");
				nRetVal = Convert.ToSingle(sValue);
			}
			else
				nRetVal = Convert.ToSingle(i);
		}
		catch { }
	}
	return nRetVal;
}
func float ToSingle(i interface{})
{
	return ToFloat(i);
}
func uint ToCount(i interface{})
{
	if nil == i)
		return 0;
	return i.ToUInt();
}
func bool ToBool(i interface{})
{
	if nil == i)
		return false;
	string sValue = i.ToString().Trim().ToLower();
	if 0 == sValue.Length || "false" == sValue)
		return false;
	if "true" == sValue)
		return true;
	try
	{
		return Convert.ToBoolean(i);
	}
	catch { }
	try
	{
		return (0 < i.ToInt32() ? true : false);
	}
	catch { }
	return false;
}


func byte[] Reverse(this byte[] aBytes, int nOffset, int nQty)
{
	IEnumerable<byte> iBytes = aBytes;
	if(0 < nOffset)
		iBytes = iBytes.Skip(nOffset);
	if(nQty < iBytes.Count())
		iBytes = iBytes.Take(nQty);
	aBytes = iBytes.Reverse().ToArray();
	return aBytes;
}
func uint ToUInt32(this byte[] aBytes, int nOffset, int nQty, bool bReverse)
{
	int nSize = sizeof(uint);
	byte[] aBuffer = aBytes;
	if nSize > nQty)
	{
		aBuffer = new byte[nSize];
		Array.Copy(aBytes, nOffset, aBuffer, nSize - nQty, nQty);
		nOffset = 0;
	}
	if bReverse)
		return BitConverter.ToUInt32(aBuffer.Reverse(nOffset, nSize), 0);
	return BitConverter.ToUInt32(aBuffer, nOffset);
}
func uint ToUInt32(this byte[] aBytes, int nOffset, bool bReverse)
{
	return aBytes.ToUInt32(nOffset, sizeof(uint), bReverse);
}
func uint ToUInt32(this byte[] aBytes, int nOffset)
{
	return aBytes.ToUInt32(nOffset, sizeof(uint), false);
}
func uint ToUInt32(this byte[] aBytes, bool bReverse)
{
	return aBytes.ToUInt32(0, bReverse);
}
func uint ToUInt32(this byte[] aBytes)
{
	return aBytes.ToUInt32(0, false);
}
func int ToInt32(this byte[] aBytes, int nOffset, int nQty, bool bReverse)
{
	int nSize = sizeof(int);
	byte[] aBuffer = aBytes;
	if nSize > nQty)
	{
		aBuffer = new byte[nSize];
		Array.Copy(aBytes, nOffset, aBuffer, nSize - nQty, nQty);
		nOffset = 0;
	}
	if bReverse)
		return BitConverter.ToInt32(aBuffer.Reverse(nOffset, nSize), 0);
	return BitConverter.ToInt32(aBuffer, nOffset);
}
func int ToInt32(this byte[] aBytes, int nOffset, bool bReverse)
{
	return aBytes.ToInt32(nOffset, sizeof(int), bReverse);
}
func int ToInt32(this byte[] aBytes, int nOffset)
{
	return aBytes.ToInt32(nOffset, sizeof(int), false);
}
func int ToInt32(this byte[] aBytes, bool bReverse)
{
	return aBytes.ToInt32(0, bReverse);
}
func int ToInt32(this byte[] aBytes)
{
	return aBytes.ToInt32(0, false);
}
func ushort ToUInt16(this byte[] aBytes, int nOffset, int nQty, bool bReverse)
{
	int nSize = sizeof(ushort);
	byte[] aBuffer = aBytes;
	if(nSize > nQty)
	{
		aBuffer = new byte[nSize];
		Array.Copy(aBytes, nOffset, aBuffer, nSize - nQty, nQty);
		nOffset = 0;
	}
	if bReverse)
		return BitConverter.ToUInt16(aBuffer.Reverse(nOffset, nSize), 0);
	return BitConverter.ToUInt16(aBuffer, nOffset);
}
func ushort ToUInt16(this byte[] aBytes, int nOffset, bool bReverse)
{
	return aBytes.ToUInt16(nOffset, sizeof(ushort), bReverse);
}
func ushort ToUInt16(this byte[] aBytes, int nOffset)
{
	return aBytes.ToUInt16(nOffset, sizeof(ushort), false);
}
func ushort ToUInt16(this byte[] aBytes, bool bReverse)
{
	return aBytes.ToUInt16(0, bReverse);
}
func ushort ToUInt16(this byte[] aBytes)
{
	return aBytes.ToUInt16(0, false);
}
func ulong ToUInt64(this byte[] aBytes, int nOffset, bool bReverse)
{
	if bReverse)
		return BitConverter.ToUInt64(aBytes.Reverse(nOffset, sizeof(ulong)), 0);
	return BitConverter.ToUInt64(aBytes, nOffset);
}
func ulong ToUInt64(this byte[] aBytes, bool bReverse)
{
	return aBytes.ToUInt64(0, bReverse);
}
func ulong ToUInt64(this byte[] aBytes, int nOffset)
{
	return BitConverter.ToUInt64(aBytes, nOffset);
}
func ulong ToUInt64(this byte[] aBytes)
{
	return aBytes.ToUInt64(0, false);
}
func double ToDouble(this byte[] aBytes, int nOffset, bool bReverse)
{
	if bReverse)
		return BitConverter.ToDouble(aBytes.Reverse(nOffset, sizeof(double)), 0);
	return BitConverter.ToDouble(aBytes, nOffset);
}
func double ToDouble(this byte[] aBytes, bool bReverse)
{
	return aBytes.ToDouble(0, bReverse);
}

func TimeSpan ToTS(i interface{})
{
	if i is TimeSpan)
		return (TimeSpan)i;

	TimeSpan tsRetVal = TimeSpan.MaxValue;
	if nil != i)
	{
		try
		{
			tsRetVal = (TimeSpan)i;
		}
		catch
		{
			try
			{
				string sValue = i.ToString();
				if sValue.Contains("day"))  // pgsql returns TS like '7 days 00:44:24'
					sValue = sValue.Replace(" days ", ".").Remove(" days").Replace(" day ", ".").Remove(" day");
				tsRetVal = TimeSpan.Parse(sValue);
			}
			catch { }
		}
	}
	return tsRetVal;
}
func IPAddress ToIP(i interface{})
{
	if i is IPAddress)
		return (IPAddress)i;
	IPAddress cRetVal = nil;
	if nil != i)
	{
		if !(i is string))
		{
			try
			{
				cRetVal = (IPAddress)i;
			}
			catch { }
		}
		cRetVal = cRetVal ?? IPAddress.Parse(i.ToString());
	}
	return cRetVal;
}

func T To<T>(i interface{})
{
	Type t = typeof(T);
	if t.IsEnum)
	{
		try
		{
			i = Enum.Parse(t, i.ToString().Trim(), true);
		}
		catch
		{
			if i.GetType().IsEnum)
				i = ((Enum)i).Translate(t); //теоретически мы сюда никогда не должны попасть... т.к. в этом случае будет выбран: func TEnum To<TEnum>(this Enum eValue)
			else
				i = i.Translate(t);
		}
	}
	else if t == typeof(byte))
		i = i.ToByte();
	else if t == typeof(short))
		i = i.ToInt16();
	else if t == typeof(ushort))
		i = i.ToUInt16();
	else if t == typeof(int))
		i = i.ToInt32();
	else if t == typeof(uint))
		i = i.ToUInt32();
	else if t == typeof(long))
		i = i.ToInt64();
	else if t == typeof(ulong))
		i = i.ToUInt64();
	else if t == typeof(bool))
		i = i.ToBool();
	else if t == typeof(float))
		i = i.ToFloat();
	else if t == typeof(string))
		i = i.ToStr();
	else if t == typeof(time.Time))
		i = i.ToDT();
	else if t == typeof(TimeSpan))
		i = i.ToTS();
	else if t == typeof(IPAddress))
		i = i.ToIP();
	return (T)i;
}

func TEnum To<TEnum>(this Enum eValue)
	where TEnum : struct
{
	TEnum eRetVal;
	if !Enum.TryParse<TEnum>(eValue.ToString().Trim(), true, out eRetVal))
		eRetVal = (TEnum)eValue.Translate(typeof(TEnum));
	return eRetVal;
}
static private object Translate(this Enum eValue, Type tEnumTarget)
{
	Enum eTarget = (Enum)Enum.GetValues(tEnumTarget).GetValue(0);
	decimal nSource = eValue.Max();
	decimal nTarget = eTarget.UnderlyingTypeMax();
	if nTarget < nSource)
		throw new InvalidCastException("target type maximum value is less than source maximum value [target:" + nTarget + "][source:" + nSource + "][value:" + eValue + "]");
	nSource = eValue.Min();
	nTarget = eTarget.UnderlyingTypeMin();
	if nTarget > nSource)
		throw new InvalidCastException("target type minimum value is greater than source minimum value [target:" + nTarget + "][source:" + nSource + "]");
	return Convert.ChangeType(eValue.ToDecimal(), Enum.GetUnderlyingType(tEnumTarget), nil);
}
static private object Translate(this object nValue, Type tEnumTarget)
{
	Enum eTarget = (Enum)Enum.GetValues(tEnumTarget).GetValue(0);
	decimal nSource = (decimal)nValue;
	decimal nTarget = eTarget.UnderlyingTypeMax();
	if nTarget < nSource)
		throw new InvalidCastException("target type maximum value is less than source value [target:" + nTarget + "][source:" + nSource + "]");
	nTarget = eTarget.UnderlyingTypeMin();
	if nTarget > nSource)
		throw new InvalidCastException("target type minimum value is greater than source value [target:" + nTarget + "][source:" + nSource + "]");
	return Convert.ChangeType(nSource, Enum.GetUnderlyingType(tEnumTarget), nil);
}

func decimal ToDecimal(this Enum eValue)
{
	return Convert.ToDecimal(eValue.ToNumeric());
}
func object ToNumeric(this Enum eValue)
{
	object oRetVal = eValue;
	Type t = Enum.GetUnderlyingType(eValue.GetType());
	if t == typeof(byte))
		oRetVal = (byte)oRetVal;
	else if t == typeof(sbyte))
		oRetVal = (sbyte)oRetVal;
	else if t == typeof(short))
		oRetVal = (short)oRetVal;
	else if t == typeof(ushort))
		oRetVal = (ushort)oRetVal;
	else if t == typeof(int))
		oRetVal = (int)oRetVal;
	else if t == typeof(uint))
		oRetVal = (uint)oRetVal;
	else if t == typeof(long))
		oRetVal = (long)oRetVal;
	else if t == typeof(ulong))
		oRetVal = (ulong)oRetVal;
	else
		throw new InvalidOperationException("unknown enum underlying type [" + t.ToString() + "]"); //LANG
	return oRetVal;
}
func decimal Max(this Enum eValue)
{
	return ToDecimal(Enum.GetValues(eValue.GetType()).Cast<Enum>().Max());
}
func decimal Min(this Enum eValue)
{
	return ToDecimal(Enum.GetValues(eValue.GetType()).Cast<Enum>().Min());
}
func decimal UnderlyingTypeMax(this Enum eValue)
{
	return Enum.GetUnderlyingType(eValue.GetType()).NumericTypeMax();
}
func decimal UnderlyingTypeMin(this Enum eValue)
{
	return Enum.GetUnderlyingType(eValue.GetType()).NumericTypeMin();
}

func decimal NumericTypeMax(this Type t)
{
	object oRetVal = nil;
	if t == typeof(byte))
		oRetVal = byte.MaxValue;
	else if t == typeof(sbyte))
		oRetVal = sbyte.MaxValue;
	else if t == typeof(short))
		oRetVal = short.MaxValue;
	else if t == typeof(ushort))
		oRetVal = ushort.MaxValue;
	else if t == typeof(int))
		oRetVal = int.MaxValue;
	else if t == typeof(uint))
		oRetVal = uint.MaxValue;
	else if t == typeof(long))
		oRetVal = long.MaxValue;
	else if t == typeof(ulong))
		oRetVal = ulong.MaxValue;
	else if t == typeof(decimal))
		oRetVal = decimal.MaxValue;
	else if t == typeof(float))
		oRetVal = float.MaxValue;
	else if t == typeof(double))
		oRetVal = double.MaxValue;
	else
		throw new ArgumentException("unknown type [" + t.ToString() + "]"); //LANG
	return Convert.ToDecimal(oRetVal);
}
func decimal NumericTypeMin(this Type t)
{
	object oRetVal = nil;
	if t == typeof(byte))
		oRetVal = byte.MinValue;
	else if t == typeof(sbyte))
		oRetVal = sbyte.MinValue;
	else if t == typeof(short))
		oRetVal = short.MinValue;
	else if t == typeof(ushort))
		oRetVal = ushort.MinValue;
	else if t == typeof(int))
		oRetVal = int.MinValue;
	else if t == typeof(uint))
		oRetVal = uint.MinValue;
	else if t == typeof(long))
		oRetVal = long.MinValue;
	else if t == typeof(ulong))
		oRetVal = ulong.MinValue;
	else if t == typeof(decimal))
		oRetVal = decimal.MinValue;
	else if t == typeof(float))
		oRetVal = float.MinValue;
	else if t == typeof(double))
		oRetVal = double.MinValue;
	else
		throw new ArgumentException("unknown type [" + t.ToString() + "]"); //LANG
	return Convert.ToDecimal(oRetVal);
}

func string ToBase64(this string sText)
{
	if nil == sText)
		return nil;
	return sText.ToBytes().ToBase64();
}
func string FromBase64(this string sText)
{
	if nil == sText)
		return nil;
	return Convert.FromBase64String(sText).ToStr();
}

func string ToBase64(this byte[] aBytes)
{
	if nil == aBytes)
		return nil;
	return Convert.ToBase64String(aBytes);
}

func byte[] ToSHA1(this byte[] aBytes)
{
	if nil == aBytes)
		return nil;
	return (new System.Security.Cryptography.SHA1Managed()).ComputeHash(aBytes);
}
func byte[] ToSHA1(this string sText)
{
	if nil == sText)
		return nil;
	return (new System.Security.Cryptography.SHA1Managed()).ComputeHash(sText.ToBytes());
}

func string ToStr(this byte[] aBytes)
{
	if nil == aBytes)
		return nil;
	return System.Text.Encoding.UTF8.GetString(aBytes, 0, aBytes.Length);
}
func byte[] ToBytes(this string sText)
{
	if nil == sText)
		return nil;
	return System.Text.Encoding.UTF8.GetBytes(sText);
}

func string ToPath(this string sPath)
{
	return System.IO.Path.GetFullPath(sPath).Replace("\\", "/").TrimEnd('/');
}
*/
