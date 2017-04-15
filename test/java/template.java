package {{ .Package }}.{{ .Class}}Test

public class {{ .Class }}Test {
{{range .Injects }}
  private {{.}}
{{end}}

{{range .Methods }}
  @Test
  public void test{{.}} {
  }
{{end}}
}