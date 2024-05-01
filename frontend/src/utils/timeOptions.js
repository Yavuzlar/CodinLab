
const dateIsValid = (date) => {
    return !Number.isNaN(new Date(date).getTime());
};

export const showDatetime = (date, options = {}) => {
    date = dateControl(date);
    if (!date) return null;

    const {
        y: year = true,
        m: month = true,
        d: day = true,
        h: hour = false,
        m: minute = false,
        nd: numericDay = true
    } = options;

    const newDate = new Date(date)
        .toLocaleDateString(navigator.language || "en", {
            ...(year && { year: 'numeric' }),
            ...(month && { month: numericDay ? 'numeric' : 'long' }),
            ...(day && { day: 'numeric' }),
            ...(hour && { hour: 'numeric' }),
            ...(minute && { minute: 'numeric' }),
        })
        .replace(/:/g, '.')
        .replace(/\//g, '.');

    return newDate;
};

export const getDayName = (date) => {
    const Days = [
        'Sunday',
        'Monday',
        'Tuesday',
        'Wednesday',
        'Thursday',
        'Friday',
        'Saturday',
    ];

    return Days[new Date(date).getDay()];
};

export const getClockTime = (options) => {
    const { date } = options;

    const hours = new Date(date).getHours();
    const minutes = new Date(date).getMinutes();

    return `${hours < 10 ? `0${hours}` : hours}.${minutes ? minutes < 10 ? `0${minutes}` : minutes : '00'}`
};

export const dateControl = (date) => {
    if (
        date === null ||
        date === undefined ||
        date === '' ||
        date === '0000-00-00' ||
        date === '0000-00-00 00:00:00' ||
        date === '0000-00-00T00:00:00.000Z' ||
        date === '0001-01-01T00:00:00Z' ||
        date === '1970-01-01' ||
        date === '1970-01-01 00:00:00' ||
        date === '1970-01-01T00:00:00.000Z' ||
        date === 0 ||
        !dateIsValid(date)
    )
        return null;

    return date;
};

export const howLongAgo = ({ date, short }) => {
    if (dateControl(date)) {
        const now = new Date();
        const past = new Date(date);

        const diff = now.getTime() - past.getTime();

        const seconds = Math.floor(diff / 1000);
        const minutes = Math.floor(seconds / 60);
        const hours = Math.floor(minutes / 60);

        const days = Math.floor(hours / 24);
        const months = Math.floor(days / 30);
        const years = Math.floor(months / 12);

        if (years > 0) return short ? `${years}y` : `${years == 1 ? `${years} year` : `${years} years`} ago`;
        else if (months > 0) return short ? `${months}mo` : `${months == 1 ? `${months} month` : `${months} months`} ago`;
        else if (days > 0) return short ? `${days}d` : `${days == 1 ? `${days} day` : `${days} days`} ago`;
        else if (hours > 0) return short ? `${hours}h` : `${hours == 1 ? `${hours} hour` : `${hours} hours`} ago`;
        else if (minutes > 0) return short ? `${minutes}m` : `${minutes == 1 ? `${minutes} minute` : `${minutes} minutes`} ago`;
        else if (seconds > 0) return short ? `${seconds}s` : `${seconds == 1 ? `${seconds} second` : `${seconds} seconds`} ago`;
        else return "şimdi";
    } else return null;
};

export const convertToISO = (date) => {
    if (!dateControl(date)) return null

    const ISODate = new Date(date)?.toISOString();
    return ISODate;
};

export const calculateDayDifference = (date) => {
    const today = new Date();
    const activityDate = new Date(date);

    const diffTime = Math.abs(today - activityDate);

    const difference = diffTime / (1000 * 60 * 60 * 24)
    const diffDays = Math.floor(difference);

    return diffDays;
}