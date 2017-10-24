"use strict"

let values = {}

function set(key, value)
{
    values[key] = value
}

function get(key, defaultValue)
{
    let retval = values[key]
    if(retval == null || typeof(retval) == "undefined")
        retval = defaultValue

    return retval
}

export default {
    set: set,
    get: get
}
