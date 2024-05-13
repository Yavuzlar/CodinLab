import {
  Button,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Typography,
} from "@mui/material";
import { useRouter } from "next/router";

const SingleItem = (props) => {
  const { title = "", path = null } = props;

  const router = useRouter();
  const handleLocate = (p) => {
    if (p) return () => router.replace(p);
  };

  const style = {
    borderRadius: 0,
    backgroundColor: "transparent",
  };

  return (
    <Button
      sx={style}
      variant="text"
      onClick={handleLocate(path)}
      active={router.pathname == path}>
      <Typography
        sx={{
          textTransform: "capitalize",
          opacity: router.pathname === path ? 1 : 0.6,
        }}>
        {" "}
        {title}{" "}
      </Typography>
    </Button>
  );
};

export default SingleItem;
