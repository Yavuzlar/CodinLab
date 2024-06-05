import ItemAcl from "../../acl/ItemAcl";
import SingleItem from "./SingleItem";

const NavItem = (props) => {
  let NavigationItem = <SingleItem {...props} />;

  return <ItemAcl item={{ ...props }}>{NavigationItem}</ItemAcl>;
};

export default NavItem;
