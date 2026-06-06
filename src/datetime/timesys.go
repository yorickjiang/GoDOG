package datetime

/***** CONST ***********************************/

type TimeSys uint8

const (
	TIME_SYS_NONE     TimeSys = iota // used when Time represents a time interval rather than an epoch
	TIME_SYS_TAI                     // international atomic time
	TIME_SYS_TT                      // terrestrial time
	TIME_SYS_UTC                     // coordinate universal time
	TIME_SYS_GPST                    // GPS time
	TIME_SYS_GLONASST                // GLONASS time
	TIME_SYS_BDT                     // BDS time
	TIME_SYS_GST                     // Galileo time
)

/***** FUNCTION ********************************/

func TimeSys2Str(s TimeSys) string {
	switch s {
	case TIME_SYS_NONE:
		return "NONE"
	case TIME_SYS_TAI:
		return "TAI"
	case TIME_SYS_TT:
		return "TT"
	case TIME_SYS_UTC:
		return "UTC"
	case TIME_SYS_GPST:
		return "GPST"
	case TIME_SYS_GLONASST:
		return "GLONASST"
	case TIME_SYS_BDT:
		return "BDT"
	case TIME_SYS_GST:
		return "GST"
	default:
		return ""
	}
}

/***********************************************/

func Str2TimeSys(s string) (TimeSys, bool) {
	switch s {
	case "NONE":
		return TIME_SYS_NONE, true
	case "TAI":
		return TIME_SYS_TAI, true
	case "TT":
		return TIME_SYS_TT, true
	case "UTC":
		return TIME_SYS_UTC, true
	case "GPST", "GPS":
		return TIME_SYS_GPST, true
	case "GLONASST", "GLONASS":
		return TIME_SYS_GLONASST, true
	case "BDT", "BDS":
		return TIME_SYS_BDT, true
	case "GST", "GAL":
		return TIME_SYS_GST, true
	default:
		return TIME_SYS_NONE, false
	}
}
