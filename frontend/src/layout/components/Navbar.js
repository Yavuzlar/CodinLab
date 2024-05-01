import { Box, Card, CardContent, ListSubheader } from "@mui/material";
import themeConfig from "@/configs/themeConfig";
import NavigationList from "./navigation";

const Navbar = () => {
  return (
    <Card
      sx={{
        width: '400px',
        maxHeight: '100vh',
        height: 'calc(100vh - 80px)',
        maxHeight: 'calc(100vh - 82px)',
        position: 'relative',
      }}
    >
      <CardContent sx={{ pb: 0 }}>
        <ListSubheader component="div" sx={{ borderRadius: '1.25rem 0rem 1.25rem 0rem', mb: '0.5rem', textAlign: 'center' }}>
          {themeConfig.templateName}
        </ListSubheader>
      </CardContent>

      <CardContent
        sx={{
          '&::-webkit-scrollbar': {
            width: '0px'
          },
          height: 'calc(100% - 88px)',
          position: 'relative',
          overflow: 'auto',
          pt: 0
        }}
      >
        <Box sx={{ height: 'auto' }}>
          <NavigationList />
        </Box>
      </CardContent>
    </Card>
  );
};

export default Navbar;
