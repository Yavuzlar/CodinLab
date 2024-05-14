import Translations from "src/components/Translations";
import C from "../assets/language/c.png";
import Cpp from "../assets/language/cpp.png";
import Go from "../assets/language/go.png";
import Js from "../assets/language/javascript.png";
import Python from "../assets/language/python.png";
import welcomeImage from "../assets/3d/3d-casual-life-young-people-working-at-the-desk.png";
import roadsImage from "../assets/3d/3d-casual-life-young-women-working-with-computer.png";
import labsImage from "../assets/3d/3ddBilgBeyaz.png";

const welcomeCard = {
  title: <Translations text="home.title" />,
  description: <Translations text="home.content" />,
  image: welcomeImage,
};

const languages = [
  {
    name: "C",
    image: C,
  },

  {
    name: "C++",
    image: Cpp,
  },

  {
    name: "Go",
    image: Go,
  },

  {
    name: "JavaScript",
    image: Js,
  },

  {
    name: "Python",
    image: Python,
  },
  
];

const roads = {
  title: <Translations text="home.roads.title" />,
  description: <Translations text="home.roads.content" />,
  image: roadsImage,
};

const labs = {
  title: <Translations text="home.labs.title" />,
  description: <Translations text="home.labs.content" />,
  image: labsImage,
};

export { welcomeCard, languages, roads, labs };
