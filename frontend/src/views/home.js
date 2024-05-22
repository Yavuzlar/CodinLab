import {
  Box,
  Card,
  CardContent,
  CircularProgress,
  Typography,
} from "@mui/material";
import { CircularProgressStatistics } from "src/components/circular-progress/CircularProgressStatistics";
import Translations from "src/components/Translations";

const Home = () => {
  const progresses = [
    {
      name: "Easy",
      value: 50,
      color: "#39CE19",
    },
    {
      name: "Medium",
      value: 25,
      color: "#EE7A19",
    },
    {
      name: "Hard",
      value: 45,
      color: "#DC0101",
    },
  ];
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

      <Card>
        <CardContent>
          <Typography>B</Typography>
          <CircularProgressStatistics progresses={progresses} />
        </CardContent>
      </Card>
    </div>
  );
};

export default Home;
