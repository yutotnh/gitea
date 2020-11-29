	"strconv"

	assertEqual(t, "sh <span class=\"added-code\">&#39;useradd -u $(stat -c &#34;%u&#34; .gitignore) jenkins</span>&#39;", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "sh &#3"},
		{Type: dmp.DiffDelete, Text: "4;useradd -u 111 jenkins&#34"},
		{Type: dmp.DiffInsert, Text: "9;useradd -u $(stat -c &#34;%u&#34; .gitignore) jenkins&#39"},
		{Type: dmp.DiffEqual, Text: ";"},
	}, DiffLineAdd))

	assertEqual(t, "<span class=\"x\">							&lt;h<span class=\"added-code\">4 class=</span><span class=\"added-code\">&#34;release-list-title df ac&#34;</span>&gt;</span>", diffToHTML("", []dmp.Diff{
		{Type: dmp.DiffEqual, Text: "<span class=\"x\">							&lt;h"},
		{Type: dmp.DiffInsert, Text: "4 class=&#"},
		{Type: dmp.DiffEqual, Text: "3"},
		{Type: dmp.DiffInsert, Text: "4;release-list-title df ac&#34;"},
		{Type: dmp.DiffEqual, Text: "&gt;</span>"},
	}, DiffLineAdd))
--- "\\a/a b/file b/a a/file"	` + `
+++ "\\b/a b/file b/a a/file"	` + `
 ` + `
--- "\\a/file with blanks" ` + `
	// Test max lines
	diffBuilder := &strings.Builder{}

	var diff = `diff --git a/newfile2 b/newfile2
new file mode 100644
index 0000000..6bb8f39
--- /dev/null
+++ b/newfile2
@@ -0,0 +1,35 @@
`
	diffBuilder.WriteString(diff)

	for i := 0; i < 35; i++ {
		diffBuilder.WriteString("+line" + strconv.Itoa(i) + "\n")
	}
	diff = diffBuilder.String()
	result, err := ParsePatch(20, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}
	result, err = ParsePatch(40, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if result.Files[0].IsIncomplete {
		t.Errorf("Files should not be incomplete! %v", result.Files[0])
	}
	result, err = ParsePatch(40, 5, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}

	// Test max characters
	diff = `diff --git a/newfile2 b/newfile2
new file mode 100644
index 0000000..6bb8f39
--- /dev/null
+++ b/newfile2
@@ -0,0 +1,35 @@
`
	diffBuilder.Reset()
	diffBuilder.WriteString(diff)

	for i := 0; i < 33; i++ {
		diffBuilder.WriteString("+line" + strconv.Itoa(i) + "\n")
	}
	diffBuilder.WriteString("+line33")
	for i := 0; i < 512; i++ {
		diffBuilder.WriteString("0123456789ABCDEF")
	}
	diffBuilder.WriteByte('\n')
	diffBuilder.WriteString("+line" + strconv.Itoa(34) + "\n")
	diffBuilder.WriteString("+line" + strconv.Itoa(35) + "\n")
	diff = diffBuilder.String()

	result, err = ParsePatch(20, 4096, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}
	result, err = ParsePatch(40, 4096, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))
	if err != nil {
		t.Errorf("There should not be an error: %v", err)
	}
	if !result.Files[0].IsIncomplete {
		t.Errorf("Files should be incomplete! %v", result.Files[0])
	}

	diff = `diff --git "a/README.md" "b/README.md"
	result, err = ParsePatch(setting.Git.MaxGitDiffLines, setting.Git.MaxGitDiffLineCharacters, setting.Git.MaxGitDiffFiles, strings.NewReader(diff))