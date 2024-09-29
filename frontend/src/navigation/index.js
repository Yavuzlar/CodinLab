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
import Translations from "src/components/Translations";

const navigation = [
  {
    path: "/home",
    title: <Translations text={"nav.home"} />,
    permission: "home",
  },
  {
    path: "/admin",
    title: <Translations text={"nav.admin"} />,
    permission: "admin",
  },
  {
    path: "/roads",
    title: <Translations text={"nav.roads"} />,
    permission: "roads",
  },
  {
    path: "/labs",
    title: <Translations text={"nav.labs"} />,
    permission: "labs",
  },
  // { // item with children
  //     type: "item",
  //     title: "Team",
  //     permission: "team",
  //     icon: <Groups3Icon />,
  //     children: [
  //         {
  //             type: "item",
  //             path: "/team/members",
  //             title: "Members",
  //             permission: "team-members",
  //             icon: <People />,
  //         },
  //         {
  //             type: "item",
  //             path: "/team/settings",
  //             title: "Settings",
  //             role: "admin",
  //             permission: "team-settings",
  //             icon: <Settings />
  //         }
  //     ]
  // },
];

export default navigation;
