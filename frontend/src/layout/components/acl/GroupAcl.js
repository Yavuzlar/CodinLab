// ** React Imports
import { useContext } from 'react'

// ** Component Imports
import { AbilityContext } from './Can'

const GroupAcl = props => {
  // ** Props
  const { children, item } = props

  // ** Hook
  const ability = useContext(AbilityContext)

  const canViewMenuGroup = it => {
    const hasAnyChildAcl = it.children && it.children.some(i => ability && ability.can("read", i.permission))
    if (!(it.permission)) {
      return hasAnyChildAcl
    }

    return ability && ability.can("read", item.permission) && hasAnyChildAcl
  }

  return item && canViewMenuGroup(item) ? <>{children}</> : null
}

export default GroupAcl
