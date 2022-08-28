import type { FC } from 'react'
import classes from './{{.Name}}.module.css'

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
