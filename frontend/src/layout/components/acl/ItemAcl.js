// ** React Imports
import { useContext } from 'react'

// ** Component Imports
import { AbilityContext } from './Can'


const ItemAcl = props => {
  // ** Props
  const { children, item } = props

  // ** Hook
  const ability = useContext(AbilityContext)

  return ability && ability.can("read", item.permission) ? <>{children}</> : null
}

export default ItemAcl
