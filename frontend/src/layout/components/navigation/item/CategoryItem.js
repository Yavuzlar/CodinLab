import { ListItem, ListItemIcon, ListItemText } from '@mui/material'

const CategoryItem = (props) => {

    const {
        title = "",
        icon = null
    } = props;

    return (
        <ListItem>
            {
                icon
                    ? <ListItemIcon>
                        {icon}
                    </ListItemIcon>
                    : null
            }

            <ListItemText primary={title} />
        </ListItem>
    )
}

export default CategoryItem