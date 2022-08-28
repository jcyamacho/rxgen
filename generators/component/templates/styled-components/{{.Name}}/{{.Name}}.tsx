import type { FC } from 'react'
import { Container } from './{{.Name}}.{{.StyledPostfix}}'

interface {{.Name}}Props {

}

const {{.Name}}: FC<{{.Name}}Props> = () => {
  return (
    <Container data-testid="{{.Name | ToKebabCase}}">
      {{.Name}}
    </Container>
  )
}

export default {{.Name}}
