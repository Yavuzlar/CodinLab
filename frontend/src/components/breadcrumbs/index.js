import { Breadcrumbs, Link, Typography } from "@mui/material";

const CustomBreadcrumbs = ({ titles }) => {
  let _titles = titles.find((e) => e.path === "/home")
    ? [...titles]
    : [
        {
          path: "/home",
          title: "Home",
          permission: "home",
        },
        ...titles,
      ];

  return (
    <Breadcrumbs separator="/" aria-label="breadcrumb">
      {_titles?.length > 1
        ? _titles.map((item, index) =>
            index == _titles.length - 1 ? (
              <Typography
                key={index}
                variant="caption"
                sx={{
                  display: "flex",
                  alignItems: "center",
                  gap: "8px",
                  cursor: "default",
                  color: (theme) => `${theme.palette.primary.main} !important`,
                }}
              >
                {item.icon}
                {item.title}
              </Typography>
            ) : !item?.path ? (
              <Typography
                key={index}
                variant="caption"
                sx={{
                  display: "flex",
                  alignItems: "center",
                  gap: "8px",
                  cursor: "default",
                }}
              >
                {item.icon}
                {item.title}
              </Typography>
            ) : (
              <Link
                underline="hover"
                key={index}
                color="inherit"
                href={item?.path}
              >
                <Typography
                  variant="caption"
                  sx={{ display: "flex", alignItems: "center", gap: "8px" }}
                >
                  {item.icon}
                  {item.title}
                </Typography>
              </Link>
            )
          )
        : null}
    </Breadcrumbs>
  );
};

export default CustomBreadcrumbs;
