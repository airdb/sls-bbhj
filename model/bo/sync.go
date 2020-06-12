package bo

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/airdb/mina-api/model/po"
)

var tmexps = []string{
	"[（|(].*[）|)]",
	"农历",
	"公历",
	"古历",
	"阳历",
	"旧历",
	"不准确",
	"大概",
	"日期不确定",
	"不确定",
	"约",
	"X日",
	"份左右",
	"期不详。",
	"。",
	"具体.*",
	"失踪.*",
	"宋振彪",
	"宋振邦2008年",
	".*身份证日期",
	"左右",
	"号",
	"阴历",
	"某天",
	"夏.*",
	"年底",
	"腊月.*",
	"九.*",
	"八.*",
	"天已记不清楚",
	"冬月.*",
	"正.*",
	"初.*",
	"下>午.*",
	"大",
	"0{4,}",
	"生.*",
	"出.*",
	"元月.*",
	"上午.*",
	"晚上.*",
	"十.*",
	"七.*",
	"五.*",
	"一.*",
	"二.*",
	"三.*",
	"四.*",
	"~.*",
	"、.*",
	"或.*",
	"&.*",
	"失",
	"————",
	"到.*",
	"暑假",
	"是",
	"和.*",
	"《.*",
	"（.*",
	"冬天",
	"之间",
	"早上.*",
	"深秋",
	"底",
	"春.*",
	"秋.*",
	"●",
	"·",
}

var (
	start      bool
	end        bool
	detailFlag bool
	// 因为详细描述多行都是value, 所以要在外层定义
	key   string
	value string
)

// 去掉html中所有标签.
func trimHTML(src string) string {
	re := regexp.MustCompile(`\<[\S\s]+?\>`)
	//将HTML标签全转换成小写
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	re = regexp.MustCompile(`\[align=left\]`)
	// src = re.ReplaceAllString(src, "\n\\[align=left\\]")
	src = re.ReplaceAllString(src, "\n")

	// //去除STYLE
	// re = regexp.MustCompile("\\<style[\\S\\s]+?\\</style\\>")
	re = regexp.MustCompile(`\<style[\S\s]+?\\</style\>`)
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re = regexp.MustCompile(`\<script[\S\s]+?\</script\>`)
	src = re.ReplaceAllString(src, "")

	// [attach]
	// src = strings.Replace(src, "[attach]", "[attach]图片ID：", -1)
	src = strings.Replace(src, "&nbsp;", "", -1)
	src = strings.Replace(src, "· ", "\r", -1)

	// 去掉[]标签, 如 [color=#000000]
	re = regexp.MustCompile(`\[[\S\s]+?\]`)
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re = regexp.MustCompile(`\<[\S\s]+?\>`)
	// src = re.ReplaceAllString(src, "\n")
	src = re.ReplaceAllString(src, "")

	// //去除连续的换行符
	// re = regexp.MustCompile("\\s{4,}")
	// src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

func formatTime(tstr string) (tm time.Time, err error) {
	timeFormats := []string{
		"2006-1月",
		"2006-01月",
		"2006年01",
		"2006年1",
		"006-1-2",
		"06年1月2",
		"20060102",
		"2006年1月2",
		"06>年1月2日",
		"2006",
		"2006-01-02 15:04:05",
		"2006-01-0215:04:05",
		"2006-01",
		"2006-1-02",
		"2006-01-2",
		"2006-1-2",
		"2006-01-02",
		"2006-0102",
		"2006--01-02",
		"2006年1月2日",
		"2006年1月02日",
		"2006年01月2日",
		"2006年01月02日",
		"2006年01月**日",
		"2006年1月",
		"2006年",
		"2006年01月02",
		"2006年01月02月",
		"2006月01月02日",
		"2006年01年02月",
		"2006年01年02日",
		"2006年1月02"}

	for _, timeFormat := range timeFormats {
		tm, err = time.ParseInLocation(timeFormat, tstr, time.Local)
		if err == nil {
			break
		}
	}

	return tm, err
}

// nolint:funlen,gocyclo,gocognit
func parseHTML(datafrom, title, msg string) (article po.Lost) {
	article.DataFrom = datafrom
	article.Title = title
	article.Subject = title

	var err error

	infoList := strings.Split(msg, "\r")

	infoLen := 1
	if len(infoList) <= infoLen {
		infoList = strings.Split(msg, "\n")
	}

	for _, info := range infoList {
		if info == "" {
			continue
		}

		log.Println(info)
		info = strings.TrimLeft(info, "\n")

		if strings.Contains(info, "本帖最后由") {
			if strings.Contains(info, "本帖最后由") {
				start = false
			}

			if !start {
				end = false
			}

			continue
		}

		if !start && (strings.Contains(info, "：") || strings.Contains(info, ":")) {
			start = true
		}

		// 如果没有标记为开始, 就已经标记为结束了，则退出
		if !start || end {
			continue
		}

		if strings.Contains(info, "站务电话") || strings.Contains(info, "注册时间") {
			detailFlag = false
		}

		if !detailFlag {
			exp := regexp.MustCompile(`.*：|.*:`)
			// 最短匹配 key, 去除 key 的部分就是 value 或是 value 的一部分
			key = exp.FindString(info)
			value = strings.Replace(info, key, "", -1)

			// 处理key 和 value
			// 删除key中的非汉字和空格
			exp = regexp.MustCompile(`(\[.+?\])|(\^M)|：|:|(\s)|[^\p{Han}]| `)
			key = exp.ReplaceAllString(key, "")
		} else {
			value = info
		}
		// fmt.Println(key, detailFlag, value)

		tmexp := regexp.MustCompile(strings.Join(tmexps, "|"))

		switch key {
		case "寻亲类别", "类别":
			article.Category = value
		case "宝贝回家编号", "编号", "寻亲编号":
			article.Babyid = value
		case "姓名":
			article.Nickname = value
		case "性别":
			// 值为1时是男性，值为2时是女性，值为0时是未知
			if value == "女" {
				// 1 --> 2
				article.Gender = 2
			} else if value == "男" {
				// 0 -> 1
				article.Gender = 1
			}
		case "失踪时身高":
			article.Height = value
		case "失踪地点", "失踪地址", "地址":
			article.MissedAddress = value
		case "失踪者特征描述":
			article.Characters = value
		case "失踪人户籍所在地", "失踪人所在省", "籍贯":
			article.BirthedProvince = value
		case "采血情况":
		case "出生日期":
			value = tmexp.ReplaceAllString(value, "")
			value = strings.Replace(value, "—", "", -1)
			article.BirthedAt, err = formatTime(value)

			if err != nil {
				fmt.Println("出生日期", value, err)
			}
		case "失踪日期", "失踪时间":
			value = tmexp.ReplaceAllString(value, "")
			value = strings.Replace(value, "—", "", -1)
			article.MissedAt, err = formatTime(value)
			// tm = tm.Format("2006-01-02 15:04:05")

			if err != nil {
				fmt.Println("失踪日期", value, err)
			}
		case "注册时间", "站务电话":
			end = true
			detailFlag = false
		case "其他资料", "其他情况", "共同经历资料":
			detailFlag = true
			article.Details += value
		default:
		}
	}

	log.Println("Babyid:", article.Babyid, ", 数据来源:", article.DataFrom)

	return article
}

func SyncFrombbs(wg *sync.WaitGroup) {
	defer wg.Done()

	for _, preForumPost := range po.GetBBSArticles() {
		datafrom := fmt.Sprintf("https://bbs.baobeihuijia.com/thread-%d-1-1.html", preForumPost.Tid)
		msg := trimHTML(preForumPost.Message)

		article := parseHTML(datafrom, preForumPost.Subject, msg)

		if article.Babyid == "" {
			fmt.Println("update datafrom only, this babyid is null.", article.DataFrom)
			// article.ID = po.FirstOrCreateArticleDataFrom(article)
			// po.UpdateArticle(article)
			article.SyncStatus = -1
		}

		po.CreateLost(article)
	}
}

func SyncFrombbsByID(tid uint) {
	preForumPost := po.GetBBSArticleByID(tid)
	datafrom := fmt.Sprintf("https://bbs.baobeihuijia.com/thread-%d-1-1.html", preForumPost.Tid)
	msg := trimHTML(preForumPost.Message)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	article := parseHTML(datafrom, preForumPost.Subject, msg)
	log.Println(article.Babyid)
	log.Println(msg)
}
