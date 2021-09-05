const querify = (obj, quotes = true) => {
    if (obj === null) return 'null';

    if (typeof obj === 'number') {
        return obj;
    }

    if (Array.isArray(obj)) {
        const props = obj.map(value => `${querify(value)}`).join(',');
        return (quotes ? `[` : '') + props + (quotes ? `]` : '');
    }

    if (typeof obj === 'object') {
        const props = Object.keys(obj)
            .map(key => `${key}:${querify(obj[key])}`)
            .join(',');
        return (quotes ? `{` : '') + props + (quotes ? `}` : '');
    }

    return JSON.stringify(obj);
};

export default querify