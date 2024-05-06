// ** Overrides Imports
import MuiChip from './chip'
import MuiContainer from './container'
import MuiButton from './button'
import MuiCard from './card'
import MuiTextField from './textfield'
import MuiFormHelperText from './formhelpertext'
import MuiGrid from './grid'
import MuiDivider from './divider'
import MuiListItem from './listitem'
import MuiTypography from './typography'

const Overrides = (theme) => {
  const chip = MuiChip(theme)
  const container = MuiContainer(theme)
  const button = MuiButton(theme)
  const card = MuiCard(theme)
  const textfield = MuiTextField(theme)
  const formhelpertext = MuiFormHelperText(theme)
  const grid = MuiGrid(theme)
  const divider = MuiDivider(theme)
  const listitem = MuiListItem(theme)
  const typography = MuiTypography(theme)

  return Object.assign(
    chip,
    container,
    button,
    card,
    textfield,
    formhelpertext,
    grid,
    divider,
    listitem,
    typography
  )
}

export default Overrides
