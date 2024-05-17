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
      name: "In progress",
      value: "30",
      color: "#8FDDFD",
    },
    {
      name: "Completed",
      value: "60",
      color: "#0A3B7A",
    },
    {
      name: "Completed",
      value: "90",
      color: "#0A3B7A",
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

      <CircularProgressStatistics progresses={progresses} />
    </div>
  );
};

export default Home;
