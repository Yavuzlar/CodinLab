import { HelpCenter, Home, People, Settings } from "@mui/icons-material";
import Groups3Icon from '@mui/icons-material/Groups3';

/*
    @data structure
    {
        type: "item" | "category" | "divider",
        path: string,
        title: string,
        role: "user" | "admin",
        icon: ReactNode,
        children: array(item)
    }
*/

const navigation = [
    { // single item
        type: "item",
        path: "/home",
        title: "Home",
        icon: <Home />,
        permission: "home",
    },
    { // single item
        type: "item",
        path: "/sample",
        title: "Sample Page",
        permission: "sample",
        icon: <HelpCenter />
    },
    { // item with children
        type: "item",
        title: "Team",
        permission: "team",
        icon: <Groups3Icon />,
        children: [
            {
                type: "item",
                path: "/team/members",
                title: "Members",
                permission: "team-members",
                icon: <People />,
            },
            {
                type: "item",
                path: "/team/settings",
                title: "Settings",
                role: "admin",
                permission: "team-settings",
                icon: <Settings />
            }
        ]
    },
]

export default navigation;