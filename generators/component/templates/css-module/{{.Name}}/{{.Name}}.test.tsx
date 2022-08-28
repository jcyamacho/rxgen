import { render, screen } from '@testing-library/react'
import {{.Name}} from './{{.Name}}'

describe('{{.Name}}', () => {
  it('renders into screen', () => {
    render(<{{.Name}} />)
    expect(
      screen.getByTestId('{{.Name | ToKebabCase}}')
    ).toBeInTheDocument()
  })
})
