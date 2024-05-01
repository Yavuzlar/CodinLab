
const findParent = (item, activePath) => {
    let titles = []

    if (item && item.length > 0) {
        item.forEach((i) => {
            if (i.path === activePath) {
                titles.push({
                    title: i.title,
                    path: i.path,
                    icon: i.icon
                })
            } else if (i.children) {
                let childTitles = findParent(i.children, activePath)

                if (childTitles.length > 0) {
                    titles.push({
                        title: i.title,
                        path: i.path,
                        icon: i.icon
                    })
                    titles = [...titles, ...childTitles]
                }
            }
        })
    }

    return titles
}

export default findParent