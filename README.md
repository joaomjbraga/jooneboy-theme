![Jooneboy](screenshots/capa.png)

Tema com alto contraste, criado com uma paleta de cores baseada em tons terrosos. Agora com suporte completo aos recursos recentes do VS Code e integração com extensões populares.

## 🖼️ Preview

![Preview](screenshots/preview.png)

## 🎨 Paleta de Cores

### Cores Principais

- **Dourado Principal**: `#C9A858` - Badges, ênfase, bordas ativas, destaque
- **Dourado Secundário**: `#AD924B` - Funções, métodos, elementos interativos
- **Verde Floresta**: `#7A8C5E` - Namespaces, módulos
- **Verde Oliva**: `#8A9B28` - Comentários (mais legível)
- **Verde Sucesso**: `#64712C` - Adições Git, testes passados, confirmações
- **Marrom Tipos**: `#B8A074` - Tipos, parâmetros de tipo
- **Verde Propriedade**: `#8B9D6F` - Propriedades, campos

### Cores de Interface

- **Background Editor**: `#0f1014` (preto moderno otimizado)
- **Background UI**: `#0F1011` (sidebar, panel)
- **Background Escuro**: `#0A0B0C` (status bar, abas inativas)
- **Foreground**: `#E8E6E3` (branco quente)
- **Números de Linha**: `#8A8470` (melhor contraste - antes `#7A745C`)
- **Números Ativos**: `#C9A858` (destaque dourado)

### Cores Semânticas

- **Keywords**: `#C678DD` (roxo suave)
- **Strings**: `#98C379` (verde claro)
- **Numbers**: `#D19A66` (laranja suave)
- **Comments**: `#8A9B28` (verde oliva - itálico, melhor contraste)
- **Doc Comments**: `#9CAF4A` (verde mais claro para documentação)
- **Errors**: `#E06C75` (vermelho)
- **Warnings**: `#E5C07B` (amarelo)
- **Info**: `#61AFEF` (azul)


### Melhorias de Contraste

🎯 **Números de Linha**: `#7A745C` → `#8A8470` (22% mais contraste)
🎯 **Abas Inativas**: Distinção clara entre ativas e inativas
🎯 **Texto Inativo**: `#808080` → `#909090` (melhor legibilidade)
🎯 **Badges**: `#AD924B` → `#C9A858` (maior destaque)
🎯 **Terminal Bright Black**: `#48453D` → `#5A564D` (muito mais legível)

### Tokens Semânticos Expandidos

```typescript
// Novos tokens semânticos
{
  "enumMember": "#E5C07B",        // Membros de enum
  "typeParameter": "#B8A074",      // Parâmetros genéricos
  "decorator": "#E5C07B",          // Decorators (@)
  "macro": "#56B6C2",              // Macros
  "label": "#C9A858",              // Labels
  "comment.documentation": "#9CAF4A" // Comentários de doc
}
```

### Suporte Específico por Linguagem

#### 🐍 Python
- F-strings com destaque especial
- Type hints diferenciados
- Decorators em itálico dourado
- Parâmetro `self` destacado

#### ⚛️ TypeScript/JavaScript
- Decorators suportados
- Utility types (`Partial`, `Pick`, etc.)
- Template literals
- JSX: Props vs atributos diferenciados

#### 🦀 Rust
- Traits em itálico
- Lifetimes destacados
- Macros com cor específica
- Mutable references em vermelho

#### 🔷 Go
- Goroutines e channels destacados
- Error type específico
- Struct fields diferenciados
- Métodos vs funções

#### ☕ Java
- Annotations douradas
- Generics com tipo específico
- Static members destacados
- Primitive types diferenciados

### Integração com Extensões

#### GitLens
```json
"gitlens.gutterBackgroundColor": "#0f101400",
"gitlens.gutterForegroundColor": "#8A8470",
"gitlens.gutterUncommittedForegroundColor": "#AD924B",
"gitlens.trailingLineForegroundColor": "#8A8470"
```

#### Error Lens
```json
"errorLens.errorBackground": "#E06C7520",
"errorLens.warningBackground": "#E5C07B20",
"errorLens.infoBackground": "#61AFEF20",
"errorLens.hintBackground": "#64712C20"
```

#### Todo Tree
```json
"todo-tree.highlights.TODOHighlight.foreground": "#E5C07B",
"todo-tree.highlights.FIXMEHighlight.foreground": "#E06C75",
"todo-tree.highlights.NOTEHighlight.foreground": "#61AFEF",
"todo-tree.highlights.HACKHighlight.foreground": "#C678DD"
```

## ⚙️ Configuração Recomendada

Para aproveitar ao máximo o tema Jooneboy:

```json
{
  "workbench.colorTheme": "Jooneboy",
  "editor.fontFamily": "'Fira Code', 'JetBrains Mono', monospace",
  "editor.fontLigatures": true,
  "editor.fontSize": 14,
  "editor.lineHeight": 22,
  "editor.cursorBlinking": "smooth",
  "editor.cursorSmoothCaretAnimation": "on",
  "editor.bracketPairColorization.enabled": true,
  "editor.guides.bracketPairs": "active",
  "editor.inlayHints.enabled": "on",
  "editor.smoothScrolling": true,
  "workbench.list.smoothScrolling": true,
  "terminal.integrated.smoothScrolling": true
}
```


## 💡 Dicas de Uso

### Fontes Recomendadas
- **Fira Code** - Excelentes ligaduras
- **JetBrains Mono** - Ótima legibilidade
- **Cascadia Code** - Moderna e limpa
- **Victor Mono** - Itálicos cursivos elegantes

### Extensões que Combinam Perfeitamente
- GitLens (controle de versão visual)
- Error Lens (erros inline)
- Todo Tree (gerenciamento de TODOs)
- Bracket Pair Colorizer (destacar brackets)
- Indent Rainbow (guias de indentação)

## 🤝 Contribuições

Sugestões e melhorias são bem-vindas! Abra uma issue ou pull request no repositório.

### Como Contribuir
1. Fork o projeto
2. Crie sua feature branch (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona MinhaFeature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abra um Pull Request

---


## Suporte

Reporte issues em:
[https://github.com/joaomjbraga/jooneboy-theme/issues](https://github.com/joaomjbraga/jooneboy-theme/issues)

| [![João M J Braga](https://github.com/joaomjbraga.png?size=100)](https://github.com/joaomjbraga)

Se você gostou deste tema, considere deixar uma ⭐ no repositório!

## 📄 Licença

Este tema é de uso livre. Criado por João Braga.