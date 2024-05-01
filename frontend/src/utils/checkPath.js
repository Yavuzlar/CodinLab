/*
@item: path | array
*/

const checkPath = (active, item, isNav = false) => {
    let check = false

    if (!isNav && Array.isArray(item)) {
        item.forEach((i) => {
            if (i.path === active) check = true
            else if (i.children) {
                check = checkPath(active, i.children)
            }
        })
    } else {
        if (item?.path === active) check = true
        else if (item?.children && !isNav) {
            check = checkPath(active, item.children)
        }
    }

    return check
}

export default checkPath