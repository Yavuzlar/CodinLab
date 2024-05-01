import { ExpandLess, ExpandMore } from '@mui/icons-material'
import { Collapse, List, ListItemButton, ListItemIcon, ListItemText } from '@mui/material'
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react'
import checkPath from '@/utils/checkPath';
import GroupAcl from '../../acl/GroupAcl';
import NavItem from './NavItem';

const MultiItem = (props) => {
    // ** States
    const [open, setOpen] = useState(false);

    const {
        title = "",
        icon = null,
        path = null,
        inside = 0,
        children
    } = props;

    // ** Hooks
    const router = useRouter()

    // ** Functions
    const handleClick = () => setOpen(!open)
    const handleLocate = (p) => { if (p) router.replace(p) }

    useEffect(() => {
        // const checkPath = children.some((item) => item.path === router.pathname)
        if (checkPath(router.pathname, props)) setOpen(true)
        else setOpen(false)

        return () => { setOpen(false) }
    }, [router.pathname])

    return (
        <GroupAcl item={props}>
            <ListItemButton
                onClick={(e) => {
                    handleClick()
                    handleLocate(path)
                }}
                sx={{ ml: inside * 2 }}
                active={checkPath(router.pathname, props, open ? true : false)}
            >
                <ListItemIcon>
                    {icon}
                </ListItemIcon>

                <ListItemText primary={title} />

                {open ? <ExpandLess /> : <ExpandMore />}
            </ListItemButton>

            <Collapse in={open} timeout="auto" unmountOnExit>
                <List component="div" disablePadding sx={{ display: 'flex', flexDirection: 'column', gap: '0.5rem' }}>
                    {
                        children.map((item, index) => {
                            return <NavItem key={index} {...item} inside={inside + 1} />
                        })
                    }
                </List>
            </Collapse>
        </GroupAcl>
    )
}

export default MultiItem