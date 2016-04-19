package helpers

import (
	sc "strconv"
	t "time"
)

func ToDT(oValue interface{}) t.Time {
	if dt, b := oValue.(t.Time); b {
		return dt
	}
	if nil != oValue {
		if dt, err := t.Parse("2006-01-02 15:04:05.999999-07:00", oValue.(string)); nil == err {
			return dt
		}
		if dt, err := t.Parse("2006-01-02 15:04:05.999+07", oValue.(string)); nil == err {
			return dt
		}
	}
	return TimeMax
}
func ToStr(oValue interface{}) *string {
	defer func() { recover() }()
	if nil == oValue {
		return nil
	}
	if s, b := oValue.(string); b {
		return &s
	}
	if dt, b := oValue.(t.Time); b {
		s := dt.Format("2006-01-02 15:04:05.99")
		return &s
	}
	if n, b := oValue.(float64); b {
		s := sc.FormatFloat(n, 'f', -1, 64)
		return &s
	}
	if n, b := oValue.(int64); b {
		s := sc.FormatInt(n, 10)
		return &s
	}
	if n, b := oValue.(int); b {
		s := sc.FormatInt(int64(n), 10)
		return &s
	}
	/*	if n, b := oValue.(ID); b {
			s := sc.FormatInt(int64(n), 10)
			return &s
		}
	*/
	if s, b := oValue.(string); b {
		return &s
	}
	return nil
}

func ToInt64(oValue interface{}) int64 {
	defer func() { recover() }()
	if n, b := oValue.(int64); b {
		return n
	}
	if s, b := oValue.(string); b {
		if n, err := sc.ParseInt(s, 10, 64); nil == err {
			return n
		}
	}
	return Int64Max
}
func ToLong(oValue interface{}) int64 {
	return ToInt64(oValue)
}
func ToInt16(oValue interface{}) int16 {
	defer func() { recover() }()
	if n, b := oValue.(int16); b {
		return n
	}
	if s, b := oValue.(string); b {
		if n, err := sc.ParseInt(s, 10, 16); nil == err {
			return int16(n)
		}
	}
	return Int16Max
}
func ToShort(oValue interface{}) int16 {
	return ToInt16(oValue)
}
func ToByte(oValue interface{}) byte {
	defer func() { recover() }()
	if n, b := oValue.(byte); b {
		return n
	}
	if s, b := oValue.(string); b {
		if n, err := sc.ParseInt(s, 10, 8); nil == err {
			return byte(n)
		}
	}
	return 255
}

/*
func ToByte(oValue interface{}) byte
{
	if oValue is byte)
		return (byte)oValue;
	byte nRetVal = byte.MaxValue;
	if nil != oValue)
	{
		try
		{
			nRetVal = Convert.ToByte(oValue);
		}
		catch { }
	}
	return nRetVal;
}
func short ToInt16(oValue interface{})
{
	if oValue is short)
		return (short)oValue;
	short nRetVal = short.MaxValue;
	if nil != oValue)
	{
		try
		{
			nRetVal = Convert.ToInt16(oValue);
		}
		catch { }
	}
	return nRetVal;
}
func ushort ToUInt16(oValue interface{})
{
	if oValue is ushort)
		return (ushort)oValue;
	ushort nRetVal = ushort.MaxValue;
	if nil != oValue)
	{
		try
		{
			nRetVal = Convert.ToUInt16(oValue);
		}
		catch { }
	}
	return nRetVal;
}
func short ToShort(oValue interface{})
{
	return oValue.ToInt16();
}
func ushort ToUShort(oValue interface{})
{
	return oValue.ToUInt16();
}
func int ToInt32(oValue interface{})
{
	if oValue is int)
		return (int)oValue;
	int nRetVal = int.MaxValue;
	if nil != oValue)
	{
		try
		{
			nRetVal = Convert.ToInt32(oValue);
		}
		catch { }
	}
	return nRetVal;
}
func uint ToUInt32(oValue interface{})
{
	if oValue is uint)
		return (uint)oValue;
	uint nRetVal = uint.MaxValue;
	if nil != oValue)
	{
		try
		{
			nRetVal = Convert.ToUInt32(oValue);
		}
		catch
		{
			if oValue is string && ((string)oValue).StartsWith("0x"))
			{
				try
				{
					nRetVal = Convert.ToUInt32((string)oValue, 16);
				}
				catch { }
			}
		}
	}
	return nRetVal;
}
func int ToInt(oValue interface{})
{
	return oValue.ToInt32();
}
func uint ToUInt(oValue interface{})
{
	return oValue.ToUInt32();
}
func long ToInt64(oValue interface{})
{
	if oValue is long)
		return (long)oValue;
	long nRetVal = long.MaxValue;
	if nil != oValue)
	{
		try
		{
			nRetVal = Convert.ToInt64(oValue);
		}
		catch { }
	}
	return nRetVal;
}
func ulong ToUInt64(oValue interface{})
{
	if oValue is ulong)
		return (ulong)oValue;
	ulong nRetVal = ulong.MaxValue;
	if nil != oValue)
	{
		try
		{
			nRetVal = Convert.ToUInt64(oValue);
		}
		catch { }
	}
	return nRetVal;
}
func long ToLong(oValue interface{})
{
	return oValue.ToInt64();
}
func ulong ToULong(oValue interface{})
{
	return oValue.ToUInt64();
}
func double ToDouble(this object oValue)
{
	return oValue.ToDouble(nil);
}
func double ToDouble(this object oValue, ushort nDecimals)
{
	return Math.Round(oValue.ToDouble(nil), nDecimals);
}
func double ToDouble(this object oValue, IFormatProvider iFormatProvider)
{
	if oValue is double)
		return (double)oValue;
	double nRetVal = double.MaxValue;
	if nil != oValue)
	{
		try
		{
			if(oValue is byte[])
				nRetVal = ((byte[])oValue).ToDouble(0, false);
			else
				nRetVal = Convert.ToDouble(oValue, iFormatProvider);
		}
		catch { }
	}
	return nRetVal;
}
func float ToFloat(oValue interface{})
{
	if oValue is float)
		return (float)oValue;
	float nRetVal = float.MaxValue;
	if nil != oValue)
	{
		try
		{
			if oValue is string)
			{
				string sValue = (string)oValue;
				sValue = sValue.Replace(".", ",");
				nRetVal = Convert.ToSingle(sValue);
			}
			else
				nRetVal = Convert.ToSingle(oValue);
		}
		catch { }
	}
	return nRetVal;
}
func float ToSingle(oValue interface{})
{
	return ToFloat(oValue);
}
func uint ToCount(oValue interface{})
{
	if nil == oValue)
		return 0;
	return oValue.ToUInt();
}
func bool ToBool(oValue interface{})
{
	if nil == oValue)
		return false;
	string sValue = oValue.ToString().Trim().ToLower();
	if 0 == sValue.Length || "false" == sValue)
		return false;
	if "true" == sValue)
		return true;
	try
	{
		return Convert.ToBoolean(oValue);
	}
	catch { }
	try
	{
		return (0 < oValue.ToInt32() ? true : false);
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

func TimeSpan ToTS(oValue interface{})
{
	if oValue is TimeSpan)
		return (TimeSpan)oValue;

	TimeSpan tsRetVal = TimeSpan.MaxValue;
	if nil != oValue)
	{
		try
		{
			tsRetVal = (TimeSpan)oValue;
		}
		catch
		{
			try
			{
				string sValue = oValue.ToString();
				if sValue.Contains("day"))  // pgsql returns TS like '7 days 00:44:24'
					sValue = sValue.Replace(" days ", ".").Remove(" days").Replace(" day ", ".").Remove(" day");
				tsRetVal = TimeSpan.Parse(sValue);
			}
			catch { }
		}
	}
	return tsRetVal;
}
func IPAddress ToIP(oValue interface{})
{
	if oValue is IPAddress)
		return (IPAddress)oValue;
	IPAddress cRetVal = nil;
	if nil != oValue)
	{
		if !(oValue is string))
		{
			try
			{
				cRetVal = (IPAddress)oValue;
			}
			catch { }
		}
		cRetVal = cRetVal ?? IPAddress.Parse(oValue.ToString());
	}
	return cRetVal;
}

func T To<T>(oValue interface{})
{
	Type t = typeof(T);
	if t.IsEnum)
	{
		try
		{
			oValue = Enum.Parse(t, oValue.ToString().Trim(), true);
		}
		catch
		{
			if oValue.GetType().IsEnum)
				oValue = ((Enum)oValue).Translate(t); //теоретически мы сюда никогда не должны попасть... т.к. в этом случае будет выбран: func TEnum To<TEnum>(this Enum eValue)
			else
				oValue = oValue.Translate(t);
		}
	}
	else if t == typeof(byte))
		oValue = oValue.ToByte();
	else if t == typeof(short))
		oValue = oValue.ToInt16();
	else if t == typeof(ushort))
		oValue = oValue.ToUInt16();
	else if t == typeof(int))
		oValue = oValue.ToInt32();
	else if t == typeof(uint))
		oValue = oValue.ToUInt32();
	else if t == typeof(long))
		oValue = oValue.ToInt64();
	else if t == typeof(ulong))
		oValue = oValue.ToUInt64();
	else if t == typeof(bool))
		oValue = oValue.ToBool();
	else if t == typeof(float))
		oValue = oValue.ToFloat();
	else if t == typeof(string))
		oValue = oValue.ToStr();
	else if t == typeof(time.Time))
		oValue = oValue.ToDT();
	else if t == typeof(TimeSpan))
		oValue = oValue.ToTS();
	else if t == typeof(IPAddress))
		oValue = oValue.ToIP();
	return (T)oValue;
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
