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
  {
    path: "/home",
    title: "Home",
    permission: "home",
  },
  {
    path: "/admin",
    title: "Admin",
    permission: "admin",
  },
  {
    path: "/roads",
    title: "Roads",
    permission: "roads",
  },
  {
    path: "/labs",
    title: "Labs",
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
