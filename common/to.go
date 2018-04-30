package helpers

import (
	"fmt"
	"math"
	sc "strconv"
	"strings"
	t "time"

	. "github.com/ratsil/go-helpers/dbc/types"
)

//ToPDT .
func ToPDT(iValue interface{}) *t.Time {
	if nil == iValue {
		return nil
	}
	if p, b := iValue.(*t.Time); b {
		return p
	}
	if dt, b := iValue.(t.Time); b {
		return &dt
	}
	if sValue, b := iValue.(string); b {
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

//ToDT .
func ToDT(iValue interface{}) t.Time {
	p := ToPDT(iValue)
	if nil == p {
		return DTNull
	}
	return *p
}

//ToPStr .
func ToPStr(iValue interface{}, a ...interface{}) *string {
	defer func() { recover() }()
	if nil == iValue {
		return nil
	}
	if s, b := iValue.(string); b {
		return &s
	}
	if p, b := iValue.(*string); b {
		return p
	}
	if dt, b := iValue.(*t.Time); b {
		if 1 > len(a) {
			a = append(a, "2006-01-02 15:04:05.99")
		}
		s := dt.Format(a[0].(string))
		return &s
	}
	if dt, b := iValue.(t.Time); b {
		if 1 > len(a) {
			a = append(a, "2006-01-02 15:04:05.99")
		}
		s := dt.Format(a[0].(string))
		return &s
	}
	if n, b := iValue.(int); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatInt(int64(n), 10)
		}
		return &s
	}
	if n, b := iValue.(int64); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatInt(n, 10)
		}
		return &s
	}
	if n, b := iValue.(uint64); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatUint(n, 10)
		}
		return &s
	}
	if n, b := iValue.(ID); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatInt(int64(n), 10)
		}
		return &s
	}
	if n, b := iValue.(float64); b {
		var s string
		if 0 < len(a) {
			s = fmt.Sprintf(a[0].(string), n)
		} else {
			s = sc.FormatFloat(n, 'f', -1, 64)
		}
		return &s
	}
	if bb, b := iValue.(bool); b {
		s := sc.FormatBool(bb)
		return &s
	}
	return nil
}

//ToStr .
func ToStr(iValue interface{}, a ...interface{}) string {
	p := ToPStr(iValue, a...)
	if nil == p {
		return ""
	}
	return *p
}

//ToInt .
func ToInt(iValue interface{}) int {
	defer func() { recover() }()
	if n, b := iValue.(int); b {
		return n
	}
	if s, b := iValue.(string); b {
		if n, err := sc.ParseInt(s, 10, 32); nil == err {
			return int(n)
		}
	}
	return math.MaxInt32
}

//ToInt64 .
func ToInt64(iValue interface{}) int64 {
	defer func() { recover() }()
	if n, b := iValue.(int64); b {
		return n
	}
	if s, b := iValue.(string); b {
		if n, err := sc.ParseInt(s, 10, 64); nil == err {
			return n
		}
	}
	return math.MaxInt64
}

//ToLong .
func ToLong(iValue interface{}) int64 {
	return ToInt64(iValue)
}

//ToInt16 .
func ToInt16(iValue interface{}) int16 {
	defer func() { recover() }()
	if n, b := iValue.(int16); b {
		return n
	}
	if s, b := iValue.(string); b {
		if n, err := sc.ParseInt(s, 10, 16); nil == err {
			return int16(n)
		}
	}
	return math.MaxInt16
}

//ToShort .
func ToShort(iValue interface{}) int16 {
	return ToInt16(iValue)
}

//ToUInt32 .
func ToUInt32(iValue interface{}) uint32 {
	defer func() { recover() }()
	if n, b := iValue.(uint32); b {
		return n
	}
	if s, b := iValue.(string); b {
		if n, err := sc.ParseUint(s, 10, 32); nil == err {
			return uint32(n)
		}
	}
	return math.MaxUint32
}

//ToUInt .
func ToUInt(iValue interface{}) uint {
	return uint(ToUInt32(iValue))
}

//ToUInt64 .
func ToUInt64(iValue interface{}) uint64 {
	defer func() { recover() }()
	if n, b := iValue.(uint64); b {
		return n
	}
	if s, b := iValue.(string); b {
		if n, err := sc.ParseUint(s, 10, 64); nil == err {
			return n
		}
	}
	return math.MaxUint64
}

//ToULong .
func ToULong(iValue interface{}) uint64 {
	return ToUInt64(iValue)
}

//ToByte .
func ToByte(iValue interface{}) byte {
	defer func() { recover() }()
	if n, b := iValue.(byte); b {
		return n
	}
	if s, b := iValue.(string); b {
		if n, err := sc.ParseInt(s, 10, 8); nil == err {
			return byte(n)
		}
	}
	return 255
}

//ToFloat64 .
func ToFloat64(iValue interface{}, a ...interface{}) (nRetVal float64) {
	defer func() { recover() }()
	nRetVal = math.MaxFloat64
	if n, b := iValue.(float64); b {
		nRetVal = n
	} else {
		if s, b := iValue.(string); b {
			if n, err := sc.ParseFloat(s, 64); nil == err {
				nRetVal = n
			}
		}
	}
	if math.MaxFloat64 > nRetVal {
		if 0 < len(a) {
			if n, b := a[0].(int); b {
				n *= 10
				nRetVal = math.Floor(nRetVal*float64(n)) / float64(n)
			}
		}
	}
	return
}

//ToDouble .
func ToDouble(iValue interface{}, a ...interface{}) float64 {
	return ToFloat64(iValue, a...)
}

//ToID .
func ToID(iValue interface{}) ID {
	return ID(ToInt64(iValue))
}

//Round .
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

//ToBool .
func ToBool(iValue interface{}) bool {
	if nil != iValue {
		if s, b := iValue.(string); b {
			s = strings.ToLower(strings.TrimSpace(s))
			return ("t" == s || "true" == s)
		}
	}
	return false
}

//ToPBool .
func ToPBool(iValue interface{}) *bool {
	if nil == iValue {
		return nil
	}
	if p, b := iValue.(*bool); b {
		return p
	}
	if bValue, b := iValue.(bool); b {
		return &bValue
	}
	return nil
}

/*
func ToByte(iValue interface{}) byte
{
	if iValue is byte)
		return (byte)iValue;
	byte nRetVal = byte.MaxValue;
	if nil != iValue)
	{
		try
		{
			nRetVal = Convert.ToByte(iValue);
		}
		catch { }
	}
	return nRetVal;
}
func short ToInt16(iValue interface{})
{
	if iValue is short)
		return (short)iValue;
	short nRetVal = short.MaxValue;
	if nil != iValue)
	{
		try
		{
			nRetVal = Convert.ToInt16(iValue);
		}
		catch { }
	}
	return nRetVal;
}
func ushort ToUInt16(iValue interface{})
{
	if iValue is ushort)
		return (ushort)iValue;
	ushort nRetVal = ushort.MaxValue;
	if nil != iValue)
	{
		try
		{
			nRetVal = Convert.ToUInt16(iValue);
		}
		catch { }
	}
	return nRetVal;
}
func short ToShort(iValue interface{})
{
	return iValue.ToInt16();
}
func ushort ToUShort(iValue interface{})
{
	return iValue.ToUInt16();
}
func int ToInt32(iValue interface{})
{
	if iValue is int)
		return (int)iValue;
	int nRetVal = int.MaxValue;
	if nil != iValue)
	{
		try
		{
			nRetVal = Convert.ToInt32(iValue);
		}
		catch { }
	}
	return nRetVal;
}
func uint ToUInt32(iValue interface{})
{
	if iValue is uint)
		return (uint)iValue;
	uint nRetVal = uint.MaxValue;
	if nil != iValue)
	{
		try
		{
			nRetVal = Convert.ToUInt32(iValue);
		}
		catch
		{
			if iValue is string && ((string)iValue).StartsWith("0x"))
			{
				try
				{
					nRetVal = Convert.ToUInt32((string)iValue, 16);
				}
				catch { }
			}
		}
	}
	return nRetVal;
}
func int ToInt(iValue interface{})
{
	return iValue.ToInt32();
}
func uint ToUInt(iValue interface{})
{
	return iValue.ToUInt32();
}
func long ToInt64(iValue interface{})
{
	if iValue is long)
		return (long)iValue;
	long nRetVal = long.MaxValue;
	if nil != iValue)
	{
		try
		{
			nRetVal = Convert.ToInt64(iValue);
		}
		catch { }
	}
	return nRetVal;
}
func ulong ToUInt64(iValue interface{})
{
	if iValue is ulong)
		return (ulong)iValue;
	ulong nRetVal = ulong.MaxValue;
	if nil != iValue)
	{
		try
		{
			nRetVal = Convert.ToUInt64(iValue);
		}
		catch { }
	}
	return nRetVal;
}
func long ToLong(iValue interface{})
{
	return iValue.ToInt64();
}
func ulong ToULong(iValue interface{})
{
	return iValue.ToUInt64();
}
func double ToDouble(this object iValue)
{
	return iValue.ToDouble(nil);
}
func double ToDouble(this object iValue, ushort nDecimals)
{
	return Math.Round(iValue.ToDouble(nil), nDecimals);
}
func double ToDouble(this object iValue, IFormatProvider iFormatProvider)
{
	if iValue is double)
		return (double)iValue;
	double nRetVal = double.MaxValue;
	if nil != iValue)
	{
		try
		{
			if(iValue is byte[])
				nRetVal = ((byte[])iValue).ToDouble(0, false);
			else
				nRetVal = Convert.ToDouble(iValue, iFormatProvider);
		}
		catch { }
	}
	return nRetVal;
}
func float ToFloat(iValue interface{})
{
	if iValue is float)
		return (float)iValue;
	float nRetVal = float.MaxValue;
	if nil != iValue)
	{
		try
		{
			if iValue is string)
			{
				string sValue = (string)iValue;
				sValue = sValue.Replace(".", ",");
				nRetVal = Convert.ToSingle(sValue);
			}
			else
				nRetVal = Convert.ToSingle(iValue);
		}
		catch { }
	}
	return nRetVal;
}
func float ToSingle(iValue interface{})
{
	return ToFloat(iValue);
}
func uint ToCount(iValue interface{})
{
	if nil == iValue)
		return 0;
	return iValue.ToUInt();
}
func bool ToBool(iValue interface{})
{
	if nil == iValue)
		return false;
	string sValue = iValue.ToString().Trim().ToLower();
	if 0 == sValue.Length || "false" == sValue)
		return false;
	if "true" == sValue)
		return true;
	try
	{
		return Convert.ToBoolean(iValue);
	}
	catch { }
	try
	{
		return (0 < iValue.ToInt32() ? true : false);
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

func TimeSpan ToTS(iValue interface{})
{
	if iValue is TimeSpan)
		return (TimeSpan)iValue;

	TimeSpan tsRetVal = TimeSpan.MaxValue;
	if nil != iValue)
	{
		try
		{
			tsRetVal = (TimeSpan)iValue;
		}
		catch
		{
			try
			{
				string sValue = iValue.ToString();
				if sValue.Contains("day"))  // pgsql returns TS like '7 days 00:44:24'
					sValue = sValue.Replace(" days ", ".").Remove(" days").Replace(" day ", ".").Remove(" day");
				tsRetVal = TimeSpan.Parse(sValue);
			}
			catch { }
		}
	}
	return tsRetVal;
}
func IPAddress ToIP(iValue interface{})
{
	if iValue is IPAddress)
		return (IPAddress)iValue;
	IPAddress cRetVal = nil;
	if nil != iValue)
	{
		if !(iValue is string))
		{
			try
			{
				cRetVal = (IPAddress)iValue;
			}
			catch { }
		}
		cRetVal = cRetVal ?? IPAddress.Parse(iValue.ToString());
	}
	return cRetVal;
}

func T To<T>(iValue interface{})
{
	Type t = typeof(T);
	if t.IsEnum)
	{
		try
		{
			iValue = Enum.Parse(t, iValue.ToString().Trim(), true);
		}
		catch
		{
			if iValue.GetType().IsEnum)
				iValue = ((Enum)iValue).Translate(t); //теоретически мы сюда никогда не должны попасть... т.к. в этом случае будет выбран: func TEnum To<TEnum>(this Enum eValue)
			else
				iValue = iValue.Translate(t);
		}
	}
	else if t == typeof(byte))
		iValue = iValue.ToByte();
	else if t == typeof(short))
		iValue = iValue.ToInt16();
	else if t == typeof(ushort))
		iValue = iValue.ToUInt16();
	else if t == typeof(int))
		iValue = iValue.ToInt32();
	else if t == typeof(uint))
		iValue = iValue.ToUInt32();
	else if t == typeof(long))
		iValue = iValue.ToInt64();
	else if t == typeof(ulong))
		iValue = iValue.ToUInt64();
	else if t == typeof(bool))
		iValue = iValue.ToBool();
	else if t == typeof(float))
		iValue = iValue.ToFloat();
	else if t == typeof(string))
		iValue = iValue.ToStr();
	else if t == typeof(time.Time))
		iValue = iValue.ToDT();
	else if t == typeof(TimeSpan))
		iValue = iValue.ToTS();
	else if t == typeof(IPAddress))
		iValue = iValue.ToIP();
	return (T)iValue;
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
