import type { FC } from 'react'
import classes from './{{.Name}}.module.scss'

interface {{.Name}}Props {

}

const {{.Name}}: FC<{{.Name}}Props> = () => {
  return (
    <div className={classes.container} data-testid="{{.Name | ToKebabCase}}">
      {{.Name}}
    </div>
  )
}

export default {{.Name}}
