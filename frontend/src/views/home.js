import { Card, CardContent, Typography } from "@mui/material";
import Translations from "src/components/Translations";

const Home = () => {
  return (
    <div>
      <Card>
        <CardContent>
          <Typography>
            <Translations text={"home.title"} />
          </Typography>
          <Typography>
            <Translations text={"active_locale"} />
          </Typography>
          <Typography>
            <Translations text={"home.title"} />
          </Typography>
          <Typography>
            <Translations text={"home.content"} />
          </Typography>
          {/* <Image src={yavuzlarLogo} /> */}
        </CardContent>
      </Card>
      <Card sx={{ mt: "10px" }}>
        <CardContent>
          <Translations text={"roads.title"} />
          <Translations text={"roads.content"} />
        </CardContent>
      </Card>
    </div>
  );
};

export default Home;
