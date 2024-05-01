import ItemAcl from "../../acl/ItemAcl"
import CategoryItem from "./CategoryItem"
import MultiItem from "./MultiItem"
import SectionTitle from "./SectionTitle"
import SingleItem from "./SingleItem"



const NavItem = (props) => {
    let NavigationItem = <SingleItem {...props} />

    switch (props.type) {
        case 'item':
            if (props.children && props.children.length > 0) return <MultiItem {...props} />
            break

        case 'category':
            NavigationItem = <CategoryItem {...props} />
            break

        case 'divider':
            NavigationItem = <SectionTitle {...props} />
            break

        case 'default':
            if (props.children && props.children.length > 0) return <MultiItem {...props} />
            break
    }

    return <ItemAcl item={{ ...props }}>{NavigationItem}</ItemAcl>
}

export default NavItem